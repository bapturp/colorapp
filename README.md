# Colorapp

Simple go application running a webserver.

It reads the `COLOR` environment variable and responds its value when requesting `/`.

## Usage

```
docker run --rm -p 8080:8080 -e COLOR=blue ghcr.io/bapturp/colorapp:latest
curl http://localhost:8080/
```

## Dev

#### Run the application locally

```sh
COLOR=blue go run colorapp.go
```

#### Build the image

```sh
docker build -t bapturp/colorapp .
```

#### Run a container

```sh
docker run --rm -p 8080:8080 -e COLOR=blue bapturp/colorapp
```
