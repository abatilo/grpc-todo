SHELL := /bin/bash

GO_VERSION = 1.12
PROTO_CONTAINER = abatilo/protobuilder

.PHONY: help
help: ## View help information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: protobuild
protobuild: ## Builds proto building container
	docker build -t $(PROTO_CONTAINER) -f Dockerfile.proto .

.PHONY: proto
proto: protobuild ## Generates protos
	docker run -v `pwd`:/src -w /src \
		-it $(PROTO_CONTAINER) ./hack/build_protos.sh
