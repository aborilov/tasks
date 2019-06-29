package user

import (
	"context"
	"tasks/pkg/user/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getServices() (*model.MockRepository, model.Service) {
	repo := &model.MockRepository{}
	service := NewService(repo)
	return repo, service
}

func TestAdd(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	name := "test name"
	user := &model.User{Name: name}
	repoOutput := &model.User{ID: 1, Name: name}
	repo.On("Add", ctx, user).Return(repoOutput, nil)
	user, err := service.Add(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	repo.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	name := "test name"
	ID := int64(1)
	repoOutput := &model.User{ID: ID, Name: name}
	repo.On("Get", ctx, ID).Return(repoOutput, nil)
	task, err := service.Get(ctx, ID)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	repo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	name := "test name"
	user := &model.User{ID: 1, Name: name}
	repoInput := &model.User{ID: 1, Name: name}
	repoOutput := &model.User{ID: 1, Name: name}
	repo.On("Update", ctx, repoInput).Return(repoOutput, nil)
	user, err := service.Update(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	repo.AssertExpectations(t)
}

func TestList(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	name := "test name"
	repoOutput := []*model.User{&model.User{ID: 1, Name: name}}
	repo.On("List", ctx).Return(repoOutput, nil)
	tasks, err := service.List(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	repo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	name := "test name"
	user := &model.User{ID: 1, Name: name}
	repo.On("Delete", ctx, user.ID).Return(nil)
	err := service.Delete(ctx, user)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
