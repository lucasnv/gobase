package domain

import (
	"time"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/validation"
)

// Value Objects
type (
	FirstName struct {
		Value string
	}

	LastName struct {
		Value string
	}

	Email struct {
		Value string
	}

	Password struct {
		Value string
	}
)

// TODO: Tengo que hacer un value object con el tiempo y para todos los value object tengo que implemenar
// El atributo value
type User struct {
	id        valueobjects.Id
	firstName FirstName
	lastName  LastName
	email     Email
	password  Password
	createdAt time.Time
	updateAt  time.Time
}

func (u *User) Id() valueobjects.Id {
	return u.id
}

func (u *User) FirstName() FirstName {
	return u.firstName
}

func (u *User) LastName() LastName {
	return u.lastName
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) Password() Password {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func NewFirstName(value string) (FirstName, errors.App) {
	validate := validation.New()

	if err := validate.Var("first_name", value, "required,alpha,lte=50"); err != nil {
		return FirstName{}, NewUserError(INVALID_USER_ERROR, err)
	}

	return FirstName{Value: value}, nil
}

func NewLastName(value string) (LastName, errors.App) {
	validate := validation.New()

	if err := validate.Var("last_name", value, "required,alpha,lte=50"); err != nil {
		return LastName{}, NewUserError(INVALID_USER_ERROR, err)
	}

	return LastName{Value: value}, nil
}

func NewEmail(value string) (Email, errors.App) {
	validate := validation.New()

	if err := validate.Var("email", value, "required,email"); err != nil {
		return Email{}, NewUserError(INVALID_USER_ERROR, err)
	}

	return Email{Value: value}, nil
}

func NewPassword(value string) (Password, errors.App) {
	validate := validation.New()

	if err := validate.Var("password", value, "required,min=4,max=15"); err != nil {
		return Password{}, NewUserError(INVALID_USER_ERROR, err)
	}

	return Password{Value: value}, nil
}

func NewUser(id valueobjects.Id, firstName FirstName, lastName LastName, email Email, password Password) User {
	return User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
		createdAt: time.Now().UTC(),
		updateAt:  time.Now().UTC(),
	}
}
