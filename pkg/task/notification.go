package task

import (
	"context"
	"fmt"
	notificationModel "tasks/pkg/notification/model"
	"tasks/pkg/task/model"
	userModel "tasks/pkg/user/model"
)

type notificationService struct {
	service      model.Service
	notification notificationModel.Service
	user         userModel.Service
}

// NewNotificationService - is a wrapper around the main service that send notification
func NewNotificationService(service model.Service, notification notificationModel.Service, user userModel.Service) model.Service {
	return &notificationService{service: service, notification: notification, user: user}
}

func (s *notificationService) Get(ctx context.Context, id int64) (*model.Task, error) {
	return s.service.Get(ctx, id)
}

func (s *notificationService) Add(ctx context.Context, task *model.Task) (*model.Task, error) {
	return s.service.Add(ctx, task)
}

func (s *notificationService) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	task, err := s.service.Update(ctx, task)
	if err != nil {
		return nil, err
	}
	for _, userID := range task.Assigned {
		user, err := s.user.Get(ctx, userID)
		if err != nil {
			return nil, err
		}
		if err := s.notification.SendEmail(ctx, user.Email, fmt.Sprintf("task changed: %s", task)); err != nil {
			return nil, err
		}
	}
	return task, nil
}

func (s *notificationService) List(ctx context.Context) ([]*model.Task, error) {
	return s.service.List(ctx)
}

func (s *notificationService) Delete(ctx context.Context, task *model.Task) error {
	return s.service.Delete(ctx, task)
}

func (s *notificationService) Assign(ctx context.Context, task *model.Task, user ...*userModel.User) (*model.Task, error) {
	return s.service.Assign(ctx, task, user...)
}
