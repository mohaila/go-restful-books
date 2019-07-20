#!/bin/sh
dep ensure -update
case $BUILD_CONFIG in
    RELEASE) 
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-s -w -extldflags "-static"'
        ;;
    *) 
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'
        ;;    
esac
