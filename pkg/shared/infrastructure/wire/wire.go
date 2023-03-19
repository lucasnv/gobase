package wire

import (
	"os"

	"<MODULE_URL_REPLACE>/pkg/users/application/deleteuser"
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
	"<MODULE_URL_REPLACE>/pkg/users/application/findusers"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
	"<MODULE_URL_REPLACE>/pkg/users/application/updateuser"
	"<MODULE_URL_REPLACE>/pkg/users/infrastructure"
)

type Wire struct {
	// Bounded Context User
	RegisterUserCommandHandler registeruser.CommandHandler
	FindUserCommandHandler     finduser.CommandHandler
	FindUsersCommandHandler    findusers.CommandHandler
	DeleteUserCommandHandler   deleteuser.CommandHandler
	UpdateUserCommandHandler   updateuser.CommandHandler
}

func Setup() *Wire {
	// Repositories
	//userRepository := infrastructure.NewInmemoryUsersRepository()
	userRepository := infrastructure.NewMongoUsersRepository(os.Getenv("MONGO_DB"))

	// Register User
	registerUserService := registeruser.NewService(userRepository)
	registerUserCommandHandler := registeruser.NewCommandHandler(*registerUserService)

	// Find User
	findUserService := finduser.NewService(userRepository)
	findUserCommandHandler := finduser.NewCommandHandler(*findUserService)

	// Find Users
	findUsersService := findusers.NewService(userRepository)
	findUsersCommandHandler := findusers.NewCommandHandler(*findUsersService)

	// Delete User
	deleteUserService := deleteuser.NewService(userRepository)
	deleteUserCommandHandler := deleteuser.NewCommandHandler(*deleteUserService)

	// Update User
	updateUserService := updateuser.NewService(userRepository)
	updateUserCommandHandler := updateuser.NewCommandHandler(*updateUserService)

	return &Wire{
		RegisterUserCommandHandler: *registerUserCommandHandler,
		FindUserCommandHandler:     *findUserCommandHandler,
		FindUsersCommandHandler:    *findUsersCommandHandler,
		DeleteUserCommandHandler:   *deleteUserCommandHandler,
		UpdateUserCommandHandler:   *updateUserCommandHandler,
	}

}
