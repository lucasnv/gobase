package findusers

import "<MODULE_URL_REPLACE>/pkg/users/domain"

type UsersResponse []userResponse

type userResponse struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt string
}

func NewUsersResponse(users domain.Users) UsersResponse {
	var response UsersResponse

	for _, u := range users.Data() {
		id := u.Id()
		createdAt := u.CreatedAt()

		response = append(response, userResponse{
			Id:        id.ToString(),
			FirstName: u.FirstName().Value,
			LastName:  u.LastName().Value,
			Email:     u.Email().Value,
			CreatedAt: createdAt.ToString(),
		})

	}

	return response
}
