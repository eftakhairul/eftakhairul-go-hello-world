# go-hello-world
A GO hello-world app docker image for testing container, kubernate and other cloud infrastucture. 

The app run on port: `8080`. It exposes following endpoints:
 - `/`  returns {"okay": true} with HTTP 200
 - `/health` returns only HTTP 200
 - `/health/ready` returns only HTTP 200
 - `/health/live` returns only HTTP 200


Run me: `docker run --rm eftakhairul/go-hello-world`

In local, build me: `docker build --no-cache -t eftakhairul/go-hello-world:latest`
