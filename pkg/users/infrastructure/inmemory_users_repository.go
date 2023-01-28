package infrastructure

import (
	"fmt"
	//domain "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type InmemoryUsersRepository struct {
	//users //array de usuarios.
}

/*
func NewInmemoryUsersRepository() InmemoryUsersRepository {
	return InmemoryUsersRepository{}
}
*/

func (r InmemoryUsersRepository) Save() error {
	fmt.Println("Save on inmemory repo")
	return nil
}

// _ userD.UserRepository = (*InmemoryUsersRepository)(nil)
