package users

type UserRepository interface {
	Find() User
	FindBy()
	Save(user User) error
}
