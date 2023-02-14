package finduser

import "<MODULE_URL_REPLACE>/pkg/users/domain"

type UserReponse struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
}

// TODO: REtornar id de usuario
func NewUserResponse(u domain.User) *UserReponse {
	return &UserReponse{
		//Id:        u.Id().Value(),
		FirstName: u.FirstName().Value,
		LastName:  u.LastName().Value,
		Email:     u.Email().Value,
	}
}
