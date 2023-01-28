package infrastructure

import "fmt"

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
func (r MongoUsersRepository) Save() error {
	fmt.Println("Save in mongo repo")
	return nil
}

//_ UserRepository = (*MongoUsersRepository)(nil)
