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

## Kubernetes
Deployment and service in kube.yml
The deployment points to a local docker registry as outlined below.

Apply to cluster with `kubectl apply -f kube.yml`

## Concourse
The docker-image-resource is pointing to a local docker registry that's in the cluster.
This needs to be applied from the `local-registry.yml` file and minikube needs to be started with `--insecure-registry localhost:5000`