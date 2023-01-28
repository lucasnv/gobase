package domain

type UserRepository interface {
	/*
		Find() User
		FindBy()
	*/
	Save() error
}
