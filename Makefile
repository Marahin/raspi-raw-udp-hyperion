APP ?= rruh
TAG ?= $(shell git rev-parse --short HEAD)
PWD ?= $(pwd)
REGISTRY ?= registry.marahin.pl

.PHONY: install-dependencies

install-dependencies:
	go mod download

docker:
	docker build -t "${REGISTRY}/${APP}:${TAG}" --build-arg APP_NAME=${APP} -f Dockerfile.prod .

push-to-registry:
	docker push "${REGISTRY}/${APP}:${TAG}"
	
	echo "${REGISTRY}/${APP}:${TAG}"

build: install-dependencies
	go build -ldflags="-w -s" -o rruh .
	chmod +x rruh

xbuild: 
	docker run --rm -v "${PWD}":/usr/src/${APP} --platform linux/arm/v6 -w /usr/src/${APP} ws2811-builder:latest go build -o "${APP}-armv6" -v