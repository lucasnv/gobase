package findusers

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
	"<MODULE_URL_REPLACE>/pkg/users/domain"
)

type userResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func NewUsersResponse(c collection.Collection) collection.Collection {

	response := c.Transform(func(item any) any {
		user := item.(domain.User)
		id := user.GetId()
		createdAt := user.GetCreatedAt()

		return userResponse{
			Id:        id.ToString(),
			FirstName: user.GetFirstName().Value,
			LastName:  user.GetLastName().Value,
			Email:     user.GetEmail().Value,
			CreatedAt: createdAt.ToString(),
		}
	})

	return response
}
