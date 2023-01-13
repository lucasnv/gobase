package registeruser

import (
	"context"

	"[REPO_URL]/pkg/shared/domain/valueobjects"
	users "[REPO_URL]/pkg/users/domain"
)

type Service struct {
}

func NewService() Service {
	return Service{}
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
