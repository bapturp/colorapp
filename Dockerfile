# Build stage
FROM node:23 AS builder
WORKDIR /usr/src/colorapp
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn build

# Final stage
FROM node:slim AS production
WORKDIR /usr/src/colorapp
RUN adduser --system colorapp --group
COPY package.json yarn.lock ./
RUN yarn install --production --frozen-lockfile
COPY --from=builder /usr/src/colorapp/dist .
USER colorapp
EXPOSE 8080
ENTRYPOINT ["node", "/usr/src/colorapp/index.js"]
