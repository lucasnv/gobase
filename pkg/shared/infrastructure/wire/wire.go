package wire

import (
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
	"<MODULE_URL_REPLACE>/pkg/users/infrastructure"
)

type Wire struct {
	// Bounded Context User
	RegisterUserCommandHandler registeruser.CommandHandler
	FindUserCommandHandler     finduser.CommandHandler
}

func Setup() *Wire {
	userRepository := infrastructure.NewInmemoryUsersRepository()

	// Register User
	registerUserService := registeruser.NewService(userRepository)
	registerUserCommandHandler := registeruser.NewCommandHandler(*registerUserService)

	// Find User
	findUserService := finduser.NewService(userRepository)
	findUserCommandHandler := finduser.NewCommandHandler(*findUserService)

	return &Wire{
		RegisterUserCommandHandler: *registerUserCommandHandler,
		FindUserCommandHandler:     *findUserCommandHandler,
	}

}
