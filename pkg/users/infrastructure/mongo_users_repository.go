package infrastructure

import (
	"fmt"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type MongoUsersRepository struct {
	//users //array de usuarios.
}

func NewMongoUsersRepository() MongoUsersRepository {
	return MongoUsersRepository{}
}

func (r *MongoUsersRepository) Save(u user.User) *errors.AppError {
	fmt.Println("Save in mongo repo")
	return nil
}

func (r *MongoUsersRepository) Find(id vo.Id) (*user.User, *errors.AppError) {
	return &user.User{}, nil
}

func (r *MongoUsersRepository) FindBy(c criteria.Criteria) (*user.Users, *errors.AppError) {

	return nil, nil
}

func (r *MongoUsersRepository) Delete(id vo.Id) *errors.AppError {
	return nil
}

var _ user.UserRepository = (*MongoUsersRepository)(nil)
