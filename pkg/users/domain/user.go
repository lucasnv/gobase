package domain

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
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
)

type User struct {
	Id        vo.Id
	FirstName FirstName
	LastName  LastName
	Email     Email
	CreatedAt vo.DateTime
}

type List []User

func (u *User) GetId() vo.Id {
	return u.Id
}

func (u *User) GetFirstName() FirstName {
	return u.FirstName
}

func (u *User) GetLastName() LastName {
	return u.LastName
}

func (u *User) GetEmail() Email {
	return u.Email
}

func (u *User) GetCreatedAt() vo.DateTime {
	return u.CreatedAt
}

func (u *User) UpdateProfile(fn FirstName, ln LastName, e Email) User {

	u.FirstName = fn
	u.LastName = ln
	u.Email = e

	return *u
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

func NewUser(id vo.Id, firstName FirstName, lastName LastName, email Email) User {
	return User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: vo.NewDateTimeNow(),
	}
}
