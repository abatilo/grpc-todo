#!/bin/bash
set -e

source ./hack/install_tools.sh
protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:pkg/api/v1/ api/proto/v1/todo/*.proto
