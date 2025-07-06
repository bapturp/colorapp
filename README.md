# Colorapp

Simple go application running a webserver. It reads the `COLOR` environment variable and respond its value when requesting `/`.

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
