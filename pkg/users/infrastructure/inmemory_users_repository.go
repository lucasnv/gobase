package infrastructure

import (
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type InmemoryUsersRepository struct {
	users []user.User
}

func NewInmemoryUsersRepository() *InmemoryUsersRepository {
	return &InmemoryUsersRepository{}
}

func (r *InmemoryUsersRepository) Save(u user.User) error {
	r.users = append(r.users, u)

	return nil
}

func (r *InmemoryUsersRepository) Find(id vo.Id) user.User {
	return r.users[0]
}

var _ user.UserRepository = (*InmemoryUsersRepository)(nil)
