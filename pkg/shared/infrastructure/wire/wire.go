package wire

import (
	"<MODULE_URL_REPLACE>/pkg/users/application/deleteuser"
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
	"<MODULE_URL_REPLACE>/pkg/users/infrastructure"
)

type Wire struct {
	// Bounded Context User
	RegisterUserCommandHandler registeruser.CommandHandler
	FindUserCommandHandler     finduser.CommandHandler
	DeleteUserCommandHandler   deleteuser.CommandHandler
}

func Setup() *Wire {
	userRepository := infrastructure.NewInmemoryUsersRepository()

	// Register User
	registerUserService := registeruser.NewService(userRepository)
	registerUserCommandHandler := registeruser.NewCommandHandler(*registerUserService)

	// Find User
	findUserService := finduser.NewService(userRepository)
	findUserCommandHandler := finduser.NewCommandHandler(*findUserService)

	// Delete User
	deleteUserService := deleteuser.NewService(userRepository)
	deleteUserCommandHandler := deleteuser.NewCommandHandler(*deleteUserService)

	return &Wire{
		RegisterUserCommandHandler: *registerUserCommandHandler,
		FindUserCommandHandler:     *findUserCommandHandler,
		DeleteUserCommandHandler:   *deleteUserCommandHandler,
	}

}
