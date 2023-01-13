# cross parameters
SHELL:=/bin/bash -O extglob
BINARY=app
VERSION=0.0.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"
CONTAINER_NAME=${cn}
MODULE_URL=${mu}

# Build step, generates the binary.
.PHONY: build
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/main.go

# Download the go lint. Not running anything.
.PHONY: lint-prepare
lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

# Run the lint across all the project. See more options https://raw.githubusercontent.com/golangci/golangci-lint .
.PHONY: lint
lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

# Run the test for all the directories.
.PHONY: test
test:
	docker-compose run --rm ${CONTAINER_NAME} go test -coverprofile test.out -v ./...

.PHONY: show-cover
show-cover:
	@clear
	docker-compose run --rm ${CONTAINER_NAME} go tool cover -html=test.out

# RUN ex: make install mu='github.com/omi-tech/api' cn="toolboard-api"
.PHONY: install
install:
	./bin/install.sh ${MODULE_URL} ${CONTAINER_NAME}
	@echo "Project base installed"

.PHONY: init
init:
	@echo ${MODULE_URL}
	docker-compose run --rm ${CONTAINER_NAME} go mod init "${MODULE_URL}"
	@echo "Project base initialized"

.PHONY: start
start:
	docker-compose up

.PHONY: stop
stop:
	docker-compose down

.PHONY: img-build
img-build:
	docker-compose build	

.PHONY: in
in:
	docker exec -it ${CONTAINER_NAME} sh

# Run go formatter
.PHONY: fmt
fmt:
	docker-compose run ${CONTAINER_NAME} gofmt -w .