# Simple REST server with Go using go-restful 
I'm using dep to manage project dependencies<br>
The server is composed of 3 files:
- model.go: for the models (Book)
- store.go: for persisting the objects in the database
- handler.go: for handling http requests for REST
- main.go: for starting the HTTP server and creating different objects needed by the server

## Build Golang build Docker image
To build the project under Linux if you are using MacOS (tested) or Windows (untested), build a Docker image for building your Go project
```bash
bash
$ docker build -t golang-1.12:build -f Dockerfile.build .
```

## Build your project
Use the previous Docker image to build your project
```bash
bash
$ docker run -v $PWD:/gobuild/src/app  golang-1.12:build
```
By default, the build configuration is Debug, if you want a release build
```bash
bash
$ docker run -e BUILD_CONFIG=RELEASE -v $PWD:/gobuild/src/app golang-1.12:build
```
## Build your Docker deployment image
```bash
bash
$ docker build -t booksapi:latest -f Dockerfile.deploy .
$ docker image ls
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
booksapi            latest              857d4f1c0967        7 seconds ago       6.99MB
golang-1.12         build               0ab52effec71        2 minutes ago       416MB
golang              1.12-alpine         6b21b4c6e7a3        8 days ago          350MB
```
