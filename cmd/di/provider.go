package di

import (
	"sync"

	"github.com/google/wire"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
	"<MODULE_URL_REPLACE>/pkg/users/domain"
	"<MODULE_URL_REPLACE>/pkg/users/infrastructure"
)

var (
	ur                 infrastructure.InmemoryUsersRepository
	userRepositoryOnce sync.Once

	rus              registeruser.Service
	registerUserOnce sync.Once

	ruch                           registeruser.CommandHandler
	registerUserCommandHandlerOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		NewInmemoryUsersRepository,
		NewRegisterUserService,
		NewRegisterUserCommandHandler,

		// bind each one of the interfaces
		wire.Bind(new(commandbus.Handler), new(registeruser.CommandHandler)),
		//wire.Bind(new(domain.UserService), new(*service)),
		wire.Bind(new(domain.UserRepository), new(infrastructure.InmemoryUsersRepository)),
	)
)

/*
func ProvideHandler(svc domain.UserService) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			svc: svc,
		}
	})

	return hdl
}*/
// Todo ver si tengo que devolver la interface en lugar de  infrastructure.InmemoryUsersRepository

func NewInmemoryUsersRepository() infrastructure.InmemoryUsersRepository {
	userRepositoryOnce.Do(func() {
		ur = infrastructure.InmemoryUsersRepository{}
	})

	return ur
}

func NewRegisterUserService(ur domain.UserRepository) registeruser.Service {
	registerUserOnce.Do(func() {
		rus = registeruser.Service{
			UserRepository: ur,
		}
	})

	return rus
}

func NewRegisterUserCommandHandler(s registeruser.Service) registeruser.CommandHandler {
	registerUserCommandHandlerOnce.Do(func() {
		ruch = registeruser.CommandHandler{
			Service: s,
		}
	})

	return ruch
}
