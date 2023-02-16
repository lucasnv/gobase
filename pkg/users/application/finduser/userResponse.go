package finduser

import "<MODULE_URL_REPLACE>/pkg/users/domain"

type UserReponse struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt string
}

func NewUserResponse(u domain.User) *UserReponse {
	id := u.Id()
	createdAt := u.CreatedAt()

	return &UserReponse{
		Id:        id.ToString(),
		FirstName: u.FirstName().Value,
		LastName:  u.LastName().Value,
		Email:     u.Email().Value,
		CreatedAt: createdAt.ToString(),
	}
}
