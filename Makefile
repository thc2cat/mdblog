
NAME= $(notdir $(shell pwd))
TAG=$(shell git  describe --tags --abbrev=0 )
FLAGS= -ldflags '-w -s -X main.Version=${NAME}-${TAG}'
BUILD= go build ${FLAGS}

blog: main.go
	CGO_ENABLED=0  ${BUILD} -a -installsuffix cgo -o blog .
