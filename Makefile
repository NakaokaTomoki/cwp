PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.0
NAME := cwp
DIST := $(NAME)-$(VERSION)

all: test cwp

build: test
	go build -o cwp $(PACKAGE_LIST)

test:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)

cwp:
	go build -o cwp $(PACKAGE_LIST)

clean:
	rm -f cwp
