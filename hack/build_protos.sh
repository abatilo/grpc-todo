#!/bin/sh

grep _ tools/tools.go | awk -F'"' '{print $2}' | xargs -tI % go install %
protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:pkg/api/v1/ api/proto/v1/foo/*.proto
