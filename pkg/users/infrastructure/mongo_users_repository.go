package infrastructure

import (
	"fmt"

	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type MongoUsersRepository struct {
	//users //array de usuarios.
}

func NewMongoUsersRepository() MongoUsersRepository {
	return MongoUsersRepository{}
}

/*
func (r MongoUsersRepository) Find() User {

}

func (r MongoUsersRepository) FindBy() {

}
*/
func (r MongoUsersRepository) Save(u user.User) error {
	fmt.Println("Save in mongo repo")
	return nil
}

func (r *MongoUsersRepository) Find(id vo.Id) user.User {
	return user.User{}
}

var _ user.UserRepository = (*MongoUsersRepository)(nil)
