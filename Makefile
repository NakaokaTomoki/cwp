PACKAGE_LIST := $(shell go list ./...)

all: test cwp

cwp:
	go build -o cwp $(PACKAGE_LIST)

test:
	go test $(PACKAGE_LIST)

clean:
	rm -f cwp
