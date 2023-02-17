package infrastructure

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type InmemoryUsersRepository struct {
	users user.UserList
}

func NewInmemoryUsersRepository() *InmemoryUsersRepository {
	return &InmemoryUsersRepository{}
}

func (r *InmemoryUsersRepository) Save(u user.User) *errors.AppError {
	r.users = append(r.users, u)

	return nil
}

func (r *InmemoryUsersRepository) Find(searchedId vo.Id) (*user.User, *errors.AppError) {

	for _, u := range r.users {
		id := u.Id()

		if id.ToString() == searchedId.ToString() {
			return &u, nil
		}
	}

	return &user.User{}, user.NewUserError(user.NOT_FOUND_ERROR)
}

func (r *InmemoryUsersRepository) FindBy(c criteria.Criteria) (*user.Users, *errors.AppError) {

	return nil, nil
}

func (r *InmemoryUsersRepository) Delete(searchedId vo.Id) *errors.AppError {

	for i, u := range r.users {
		id := u.Id()

		if id.ToString() == searchedId.ToString() {
			r.users = append(r.users[:i], r.users[i+1:]...)
		}
	}

	return nil
}

var _ user.UserRepository = (*InmemoryUsersRepository)(nil)
