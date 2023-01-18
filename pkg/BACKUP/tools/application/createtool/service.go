package createtool

import (
	"context"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) exec(ctx context.Context, id, name, link, description string) error {
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
