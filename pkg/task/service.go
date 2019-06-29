package task

import (
	"context"
	"tasks/pkg/task/model"
	userModel "tasks/pkg/user/model"
)

type service struct {
	repo model.Repository
}

// NewService - return task service implementation
func NewService(repo model.Repository) model.Service {
	return &service{repo: repo}
}

func (s *service) Get(ctx context.Context, id int64) (*model.Task, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) Add(ctx context.Context, task *model.Task) (*model.Task, error) {
	task.Status = model.StatusNew
	return s.repo.Add(ctx, task)
}

func (s *service) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	return s.repo.Update(ctx, task)
}

func (s *service) List(ctx context.Context) ([]*model.Task, error) {
	return s.repo.List(ctx)
}

func (s *service) Delete(ctx context.Context, task *model.Task) error {
	return s.repo.Delete(ctx, task.ID)
}

func (s *service) Assign(ctx context.Context, task *model.Task, user ...*userModel.User) (*model.Task, error) {
	userIDs := make([]int64, len(user))
	for i, u := range user {
		userIDs[i] = u.ID
	}
	task.Assigned = userIDs
	return s.repo.Update(ctx, task)
}
