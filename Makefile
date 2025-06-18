REGISTRY ?= ghcr.io/deniszm
VERSION ?= $(shell git describe --tags --abbrev=0)-$(shell git rev-parse HEAD|cut -c1-7)
HOST_ARCH = $(shell uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/')

TARGETOS ?= linux
TARGETARCH ?= arm64
CGO_ENABLED ?= 0

BIN_NAME = obot
BUILD_DIR = build

.PHONY: all clean test image image-host-arch push build build-local

all: test build image


test:
	@echo "Running tests..."
	@go test -v -cover ./...

build:
	docker run --rm -v $(PWD):/src -w /src quay.io/projectquay/golang:1.24 \
	  /bin/sh -c 'GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) CGO_ENABLED=$(CGO_ENABLED) \
	    go build \
	      -ldflags "-s -w -X=github.com/deniszm/obot/cmd.appVersion=$(VERSION)" \
	      -o $(BUILD_DIR)/$(BIN_NAME)-$(TARGETOS)-$(TARGETARCH) \
	      main.go'

# Local build without Docker (for use in Dockerfile)
build-local:
	mkdir -p $(BUILD_DIR)
	GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) CGO_ENABLED=$(CGO_ENABLED) \
	  go build \
	    -ldflags "-s -w -X=github.com/deniszm/obot/cmd.appVersion=$(VERSION)" \
	    -o $(BUILD_DIR)/$(BIN_NAME)-$(TARGETOS)-$(TARGETARCH) \
	    main.go

image:
	docker build . \
	  --platform $(TARGETOS)/$(TARGETARCH) \
	  --build-arg TARGETARCH=$(TARGETARCH) \
	  --build-arg VERSION=$(VERSION) \
	  --tag $(REGISTRY)/obot:$(VERSION)-$(TARGETARCH) \
	  --load

# Default image build for current host architecture
image-host-arch:
	docker build . \
	  --build-arg TARGETARCH=$(HOST_ARCH) \
	  --build-arg VERSION=$(VERSION) \
	  --tag $(REGISTRY)/obot:$(VERSION)-$(HOST_ARCH) \
	  --load

push:
	docker push $(REGISTRY)/obot:$(VERSION)-$(TARGETARCH)

clean:
	docker rmi $(REGISTRY)/obot:$(VERSION)-$(TARGETARCH) || true
	rm -rf $(BUILD_DIR)
