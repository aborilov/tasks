package comment

import (
	"context"
	"tasks/pkg/comment/model"
	taskModel "tasks/pkg/task/model"
	"tasks/pkg/user"
	userModel "tasks/pkg/user/model"
)

type service struct {
	repo model.Repository
}

// NewService - return comment service implementation
func NewService(repo model.Repository) model.Service {
	return &service{repo: repo}
}

func (s *service) Get(ctx context.Context, id int64) (*model.Comment, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) Add(ctx context.Context, task *taskModel.Task, msg string) (*model.Comment, error) {
	user, ok := user.FromContext(ctx)
	if !ok {
		return nil, userModel.ErrUnAuthorized
	}
	return s.repo.Add(ctx, task.ID, user.ID, msg)
}

func (s *service) Update(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	return s.repo.Update(ctx, comment)
}

func (s *service) List(ctx context.Context) ([]*model.Comment, error) {
	return s.repo.List(ctx)
}

func (s *service) Delete(ctx context.Context, comment *model.Comment) error {
	user, ok := user.FromContext(ctx)
	if !ok {
		return userModel.ErrUnAuthorized
	}
	if user.Role != userModel.RoleAdmin {
		return userModel.ErrPermissionDenied
	}
	return s.repo.Delete(ctx, comment.ID)
}
