#!/bin/bash
set -e

source ./hack/install_tools.sh
go generate ./...
