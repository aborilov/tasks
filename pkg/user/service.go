package user

import (
	"context"
	"tasks/pkg/user/model"
)

type service struct {
	repo model.Repository
}

// NewService - return user service implementation
func NewService(repo model.Repository) model.Service {
	return &service{repo: repo}
}

func (s *service) Get(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) Add(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.Add(ctx, user)
}

func (s *service) Update(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.Update(ctx, user)
}

func (s *service) List(ctx context.Context) ([]*model.User, error) {
	return s.repo.List(ctx)
}

func (s *service) Delete(ctx context.Context, user *model.User) error {
	return s.repo.Delete(ctx, user.ID)
}
