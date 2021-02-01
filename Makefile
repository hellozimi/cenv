.PHONY: clean

ifndef VERSION
	VERSION := $(shell git describe --tags --always --dirty="-dev")
endif

LDFLAGS := -ldflags='-X github.com/hellozimi/cenv/pkg/version.version=$(VERSION)'

clean:
	rm -rf ./build

all: build/cenv-$(VERSION)-darwin-amd64 build/cenv-$(VERSION)-darwin-arm64 build/cenv-$(VERSION)-linux-amd64

build/:
	mkdir -p build

macos: build/cenv-$(VERSION)-darwin-amd64

linux: build/cenv-$(VERSION)-linux-amd64
	
build/cenv-$(VERSION)-darwin-amd64: | build/
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -mod=vendor $(LDFLAGS) -o $@

build/cenv-$(VERSION)-darwin-arm64: | build/
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -mod=vendor $(LDFLAGS) -o $@
	
build/cenv-$(VERSION)-linux-amd64: | build/
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -mod=vendor $(LDFLAGS) -o $@

