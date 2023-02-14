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

> VsCode pre-installed

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

### Work with Visual Studio Code and the container

- First of all you have to open a remote window and then attach to running container (you have to take into account, you will be able work only when your container will be up and running, make start)
- If is the first time that you work on this mode you going to have to install for sure a visual studio code plugin for golang into the container

## TODO

- Finish DI wire implementing
- I have to finish to create an example to create an user, and get an user. I have to implement mongo and memory storage
- move http server to share/infrastructure
- Take a look on make commands like lint-prepare and lint (make a intallation and it's not necesary)
- Resolve todo comments
- Ver todos los news si tengo que devolver un puntero

## Configurate .vscode to debug

- We have created a launch.json file in order to help you with VScode's configuration

## Good practices

- The recommended way to organize a Go file is to start with package declaration and import statements, followed by constants, variables, types, functions, and interfaces in that order.
