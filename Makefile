SHELL := /bin/bash
TOOLS_CONTAINER = abatilo/grpc-todo-tools

.PHONY: help
help: ## View help information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Clean any generated files
	rm -rf todo
	find -name *.pb.go -delete
	find -name mock_*.go -delete

.PHONY: tools
tools:
	docker build -t $(TOOLS_CONTAINER) -f Dockerfile.tools .

.PHONY: proto
proto: clean tools
	docker run -v `pwd`:/src -w /src \
		-it $(TOOLS_CONTAINER) ./hack/build_protos.sh

.PHONY: mock
mock: clean tools
	docker run -v `pwd`:/src -w /src \
		-it $(TOOLS_CONTAINER) ./hack/build_mocks.sh

.PHONY: generate
generate: proto mock ## Generate protos and mocks
	sudo chown -R $(shell whoami) .
