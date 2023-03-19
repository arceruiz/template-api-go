OS ?= linux
ARCH ?= amd64
VERSION ?= $$(git rev-parse --short HEAD)
TAG ?= "arceruiz/template-api-go:$(VERSION)"

build: 
	env GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -o ./build/package/main_$(OS)_$(ARCH) ./cmd/service/main.go 

push: bake	
	docker push $(TAG) && \
	echo "New version: $(TAG)" && \
	echo "\033[1;32mPublish was successful\033[0m"

run: bake
	docker run -v "$(pwd)"/configs/:/app/configs/ -v "$(pwd)"/configs/:/app/secrets/ $(TAG)

bake: build
	cd ./build/package && \
	docker build -t $(TAG) --build-arg GOARCH=$(ARCH) .