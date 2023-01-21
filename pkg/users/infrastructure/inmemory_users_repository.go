package users

type InmemoryUsersRepository struct {
	//users //array de usuarios. 
}

func NewInmemoryUsersRepository () InmemoryUsersRepository {
	return InmemoryUsersRepository{}
}

func (r InmemoryUsersRepository) Find() User {

}

func (r InmemoryUsersRepository) FindBy() {

}

func (r InmemoryUsersRepository) Save() error {

}

_ UserRepository = (*InmemoryUsersRepository)(nil)