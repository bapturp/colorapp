# Build stage
FROM golang:1.25.4-bookworm AS builder
WORKDIR /usr/src/colorapp
COPY . .
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/colorapp

# Final stage
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /usr/local/bin/colorapp /colorapp
EXPOSE 8080
ENTRYPOINT ["/colorapp"]
