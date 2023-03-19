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
	var response []userResponse
	users := c.Data().(domain.List)

	for _, u := range users {
		id := u.GetId()
		createdAt := u.GetCreatedAt()

		response = append(response, userResponse{
			Id:        id.ToString(),
			FirstName: u.GetFirstName().Value,
			LastName:  u.GetLastName().Value,
			Email:     u.GetEmail().Value,
			CreatedAt: createdAt.ToString(),
		})
	}

	return c.SetData(response)
}
