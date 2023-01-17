## About this project

This is GoLang base structure. It works as a good starting point to create a GoLang app from scratch in seconds.

## How to use it?

- The best and easiest way to install this project is using the go-base-installer [Installer](https://github.com/lucasnv/gobase-installer).

## Characteristics

> Hexagonal architecture implemented.

> Gin pre-installed.

> Env tool pre-installed

> Docker

> Some dev tools.

> Golang V1.19.5

## Useful commands

You can execute any command using makefile, the following command described below are available.

- make init
- make start
- make stop
- make img-build
- make in
- make test
- make fmt
- make test
- make show-cover
- make lint
- make build
- make lint-prepare

## TODO

- Delete tool boundedcontext.
- delete internal folder and move everything into share/infrastructure
- Create an example to create an user, and get an user.
- move http server to share/infrastructure
- Create a folde into cmd folder one per entrypoint (api, cli, etc)
- Take a look on make commands like lint-prepare and lint (make a intallation and it's not necesary)
