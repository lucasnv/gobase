FROM golang:1.19.5-alpine3.17 AS base

WORKDIR /opt/app/api

# System dependencies
RUN apk update \
    && apk add --no-cache \
    && apk add build-base \
    ca-certificates \
    git \
    && update-ca-certificates


### Development with hot reload and debugger
FROM base AS dev

WORKDIR /opt/app/api

### In order tu to make easear the debuger configuration using vscode
COPY launch.json /opt/app/.vscode/launch.json

# Install tools
RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

# Install go libs
RUN go install golang.org/x/tools/gopls@latest
RUN go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

# Install DI
RUN cd /opt/app/api/cmd/di && go run github.com/google/wire/cmd/wire
RUN cd /opt/app/api/

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest

# Script for init dlv debugging
COPY bin/init_debugging.sh /usr/local/bin/debug
RUN chmod +x /usr/local/bin/debug
RUN export GOFLAGS=-buildvcs=false

EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /opt/app/api

# Application dependencies
COPY . /opt/app/api
RUN go mod download \
    && go mod verify

RUN go build -o api -a .

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /opt/app/api /usr/local/bin/api
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/api"]