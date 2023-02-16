package deleteuser

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type Service struct {
	UserRepository user.UserRepository
}

func NewService(r user.UserRepository) *Service {
	return &Service{
		UserRepository: r,
	}
}

func (s *Service) exec(ctx context.Context, id valueobjects.Id) errors.App {

	err := s.UserRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
