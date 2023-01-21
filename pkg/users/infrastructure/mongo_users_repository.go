package users

type MongoUsersRepository struct {
	//users //array de usuarios. 
}

func NewMongoUsersRepository () MongoUsersRepository {
	return MongoUsersRepository{}
}

func (r MongoUsersRepository) Find() User {

}

func (r MongoUsersRepository) FindBy() {

}

func (r MongoUsersRepository) Save() error {

}

_ UserRepository = (*MongoUsersRepository)(nil)