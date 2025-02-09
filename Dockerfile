# Build stage
FROM node:22-alpine AS builder
WORKDIR /usr/src/colorapp
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn build

# Final stage
FROM node:22-alpine
WORKDIR /usr/src/colorapp
ENV USER=colorapp
RUN addgroup $USER && adduser --disabled-password --system --ingroup $USER $USER
COPY package.json yarn.lock ./
RUN yarn install --production --frozen-lockfile
COPY --from=builder /usr/src/colorapp/dist .
USER $USER
EXPOSE 8080
ENTRYPOINT ["node", "/usr/src/colorapp/index.js"]
