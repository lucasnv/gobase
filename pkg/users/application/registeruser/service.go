package registeruser

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	users "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type Service struct {
	userRepository users.UserRepository
}

func NewService(repository users.UserRepository) Service {
	return Service{
		userRepository: repository,
	}
}

func (s Service) exec(ctx context.Context, id valueobjects.Id, firstName users.FirstName, lastName users.LastName, email users.Email, password users.Password) error {
	return nil
	/*course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, course.PullEvents())*/
}
