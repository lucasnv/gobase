package wire

import (
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/criteria"
	"<MODULE_URL_REPLACE>/pkg/users/application/deleteuser"
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
	"<MODULE_URL_REPLACE>/pkg/users/application/findusers"
	"<MODULE_URL_REPLACE>
	"<MODULE_URL_REPLACE>/infrastructure"
)

type Wire struct {
	// Bounded Context User
	RegisterUserCommandHandler registeruser.CommandHandler
	FindUserCommandHandler     finduser.CommandHandler
	FindUsersCommandHandler    findusers.CommandHandler
	DeleteUserCommandHandler   deleteuser.CommandHandler
}

func Setup() *Wire {
	// Generics
	criteria, _ := criteria.NewInmemoryBuilder()

	// Repositories
	userRepository := infrastructure.NewInmemoryUsersRepository()

	// Register User
	registerUserService := registeruser.NewService(userRepository)
	registerUserCommandHandler := registeruser.NewCommandHandler(*registerUserService)

	// Find User
	findUserService := finduser.NewService(userRepository)
	findUserCommandHandler := finduser.NewCommandHandler(*findUserService)

	// Find Users
	findUsersService := findusers.NewService(userRepository)
	findUsersCommandHandler := findusers.NewCommandHandler(*findUsersService, criteria)

	// Delete User
	deleteUserService := deleteuser.NewService(userRepository)
	deleteUserCommandHandler := deleteuser.NewCommandHandler(*deleteUserService)

	return &Wire{
		RegisterUserCommandHandler: *registerUserCommandHandler,
		FindUserCommandHandler:     *findUserCommandHandler,
		FindUsersCommandHandler:    *findUsersCommandHandler,
		DeleteUserCommandHandler:   *deleteUserCommandHandler,
	}

}
