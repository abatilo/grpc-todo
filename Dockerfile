FROM golang:1.12 as backend

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ARG PROTOC_VERSION=3.11.4

# Install protoc
RUN apt-get update && apt-get install -y --no-install-recommends unzip=* \
      && apt-get clean \
      && rm -rf /var/lib/apt/lists/*

WORKDIR /tmp
RUN curl -Lo /tmp/protoc.zip "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip" \
      && unzip /tmp/protoc.zip \
      && mv /tmp/bin/* /usr/local/bin/ \
      && mv /tmp/include/* /usr/local/include/ \
      && rm -rf /tmp/*

# Install dependencies
WORKDIR /go/src/github.com/abatilo/grpc-todo
COPY ./go.mod ./go.sum ./
RUN go mod download

# Build artifacts
COPY ./hack ./hack
COPY ./tools ./tools
RUN /go/src/github.com/abatilo/grpc-todo/hack/install_tools.sh

COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./mock ./mock
COPY ./api ./api
RUN /go/src/github.com/abatilo/grpc-todo/hack/build_protos.sh \
      && /go/src/github.com/abatilo/grpc-todo/hack/build_mocks.sh \
      && go test ./... \
      && go build -ldflags="-w -s" -o /go/bin/todo cmd/todo.go

FROM scratch
# SSL Certs
COPY --from=backend /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Health probe
COPY --from=backend /go/bin/grpc-health-probe /go/bin/grpc-health-probe

# Copy our static executable
COPY --from=backend /go/bin/todo /go/bin/todo

ENTRYPOINT ["/go/bin/grpc-health-probe", "-addr=:6666"]
# ENTRYPOINT ["/go/bin/todo"]
