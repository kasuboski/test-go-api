# test-go-api
An example api in go to play with go, docker, and kubernetes

## Api
Uses `go-chi` and just returns json at /
```
{
  "greeting": "hello"
}
```

It starts on the port given by the `PORT` environment variable or 8080 if it's not set.

It can be started with `go run test-go-api.go`

## Docker
The Dockerfile is a multi-stage build.

The first stage is called build and installs the dependencies and builds the binary.

The second stage copies that binary sets the port and runs the binary.

Build it with `docker build -t kasuboski/test-go-api:latest .`

Run it with `docker run -p 8080:8080 kasuboski/test-go-api`