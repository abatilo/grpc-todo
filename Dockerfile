FROM golang:1.12 as backend

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Install dependencies
WORKDIR /go/src/github.com/abatilo/grpc-todo
COPY ./go.mod ./go.sum ./
RUN go mod download

# Build artifacts
WORKDIR /go/src/github.com/abatilo/todo

COPY ./hack ./hack
COPY ./tools ./tools
RUN /go/src/github.com/abatilo/todo/hack/install_tools.sh

COPY ./cmd ./cmd
COPY ./pkg ./pkg
RUN go build -ldflags="-w -s" -o /go/bin/todo cmd/todo.go

FROM scratch
# SSL Certs
COPY --from=backend /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Health probe
COPY --from=backend /go/bin/grpc-health-probe /go/bin/grpc-health-probe

# Copy our static executable
COPY --from=backend /go/bin/todo /go/bin/todo

ENTRYPOINT ["/go/bin/grpc-health-probe", "-addr=:6666"]
# ENTRYPOINT ["/go/bin/todo"]
