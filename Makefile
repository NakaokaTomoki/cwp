PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.0
NAME := cwp
DIST := $(NAME)-$(VERSION)

$(warning PACKAGE_LIST = $(PACKAGE_LIST))

all: build test

build:
	go build -o cwp $(PACKAGE_LIST)

test:
	gofmt -l -s .
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)

clean:
	rm -f cwp
