VERSION := $(shell git describe --tags)

all: tools fmt build lint

tools:
	go get -u github.com/golang/lint/golint

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./...

build:
	 CGO_ENABLED=0 go build -o "terrapin-`uname -s`-`uname -m`"
	 ln -sf "terrapin-`uname -s`-`uname -m`" terrapin

lint:
	golint

clean:
	rm -f ./terrapin

.PHONY: fmt build lint clean