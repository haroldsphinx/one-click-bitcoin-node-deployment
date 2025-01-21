ENVIRONMENT ?= development
NAMESPACE ?= bitcoin-exporter
IMAGE_NAME := bitcoin-exporter-go
VERSION := $(shell git rev-parse --short HEAD)
REGISTRY_URL := haroldsphinx/$(IMAGE_NAME):latest

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOARM ?= $(shell go env GOARM)
CGO_ENABLED?= $(shell go env CGO_ENABLED)
GO_ENV := GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) CGO_ENABLED=$(CGO_ENABLED)


.PHONY: build push deploy clean validate build-bitcoin-exporter

build-bitcoin-exporter:
	$(GO_ENV) go build $(GO_FLAGS) -o ./cmd/bitcoin-exporter ./main.go

build: validate
	REGISTRY_URL=$(REGISTRY_URL) \
	docker build -t ${REGISTRY_URL} .

push: build
	@echo "pushing image to ${REGISTRY_URL}"
	REGISTRY_URL=$(REGISTRY_URL) \
	docker push $(REGISTRY_URL)

clean:
	docker system prune -f
	
.PHONY: validate
validate:
ifeq (${ENVIRONMENT},)
	$(error ENVIRONMENT is undefined,  Should be development, staging or production)
endif
