REGESTRY=ghcr.io/deniszm
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse HEAD|cut -c1-7)

BIN_NAME=obot
BUILD_DIR=build

.PHONY: all image-x86 image-arm push-x86 push-arm clean
.PHONY: build-linux-x86 build-linux-arm build-darwin-x86 build-darwin-arm build-windows-x86
.PHONY: image

all: image-x86 image-arm

# Default image build for current host architecture
image:
	docker buildx build . \
	  --build-arg TARGETOS=linux \
	  --build-arg TARGETARCH=$$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/') \
	  --build-arg VERSION=$(VERSION) \
	  --tag $(REGESTRY)/obot:$(VERSION)-$$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/') \
	  --load

image-x86:
	docker buildx build . \
	  --platform linux/amd64 \
	  --build-arg TARGETOS=linux \
	  --build-arg TARGETARCH=amd64 \
	  --build-arg VERSION=$(VERSION) \
	  --tag $(REGESTRY)/obot:$(VERSION)-amd64 \
	  --load

image-arm:
	docker buildx build . \
	  --platform linux/arm64 \
	  --build-arg TARGETOS=linux \
	  --build-arg TARGETARCH=arm64 \
	  --build-arg VERSION=$(VERSION) \
	  --tag $(REGESTRY)/obot:$(VERSION)-arm64 \
	  --load

push-x86:
	docker push $(REGESTRY)/obot:$(VERSION)-amd64

push-arm:
	docker push $(REGESTRY)/obot:$(VERSION)-arm64

build-linux-x86:
	docker run --rm -v $(PWD):/src -w /src quay.io/projectquay/golang:1.24 \
	  /bin/sh -c 'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BIN_NAME)-linux-amd64 main.go'

build-linux-arm:
	docker run --rm -v $(PWD):/src -w /src quay.io/projectquay/golang:1.24 \
	  /bin/sh -c 'GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BIN_NAME)-linux-arm64 main.go'

build-darwin-x86:
	docker run --rm -v $(PWD):/src -w /src quay.io/projectquay/golang:1.24 \
	  /bin/sh -c 'GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BIN_NAME)-darwin-amd64 main.go'

build-darwin-arm:
	docker run --rm -v $(PWD):/src -w /src quay.io/projectquay/golang:1.24 \
	  /bin/sh -c 'GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BIN_NAME)-darwin-arm64 main.go'

build-windows-x86:
	docker run --rm -v $(PWD):/src -w /src quay.io/projectquay/golang:1.24 \
	  /bin/sh -c 'GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BIN_NAME)-windows-amd64.exe main.go'

clean:
	docker rmi $(REGESTRY)/obot:$(VERSION)-amd64 || true
	docker rmi $(REGESTRY)/obot:$(VERSION)-arm64 || true
	rm -rf $(BUILD_DIR)

