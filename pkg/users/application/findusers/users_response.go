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

	return c.SetData(response)
}
