# Pull base image
FROM golang:1.17.6 as base

FROM base as dev

# Install git
RUN apt update && apt install git

RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api

CMD ["air"]