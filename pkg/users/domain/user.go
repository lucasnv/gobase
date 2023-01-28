package domain

import (
	"time"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/validation"
)

// Value Objects
type (
	FirstName struct {
		value string
	}

	LastName struct {
		value string
	}

	Email struct {
		value string
	}

	Password struct {
		value string
	}
)

func NewFirstName(value string) (FirstName, error) {
	validate := validation.New()

	if err := validate.Var("first_name", value, "required,alpha,lte=50"); err != nil {
		return FirstName{}, err
	}

	return FirstName{value: value}, nil
}

func NewLastName(value string) (LastName, error) {
	validate := validation.New()

	if err := validate.Var("last_name", value, "required,alpha,lte=50"); err != nil {
		return LastName{}, err
	}

	return LastName{value: value}, nil
}

func NewEmail(value string) (Email, error) {
	validate := validation.New()

	if err := validate.Var("email", value, "required,email"); err != nil {
		return Email{}, err
	}

	return Email{value: value}, nil
}

func NewPassword(value string) (Password, error) {
	validate := validation.New()

	err := validate.Var("password", value, "required,gte=8")

	if value == "" {
		return Password{}, err
	}

	return Password{value: value}, nil
}

// Entity
type User struct {
	id        valueobjects.Id
	firstName FirstName
	lastName  LastName
	email     Email
	password  Password
	createdAt time.Time
	updateAt  time.Time
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
