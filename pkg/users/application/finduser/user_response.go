package finduser

import "<MODULE_URL_REPLACE>/pkg/users/domain"

type UserResponse struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt string
}

func NewUserResponse(u domain.User) UserResponse {
	id := u.GetId()
	createdAt := u.GetCreatedAt()

	return UserResponse{
		Id:        id.ToString(),
		FirstName: u.GetFirstName().Value,
		LastName:  u.GetLastName().Value,
		Email:     u.GetEmail().Value,
		CreatedAt: createdAt.ToString(),
	}
}
