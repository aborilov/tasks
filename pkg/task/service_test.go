package task

import (
	"context"
	"tasks/pkg/task/model"
	userModel "tasks/pkg/user/model"
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
	desc := "some test task"
	task := &model.Task{Description: desc}
	repoInput := &model.Task{Description: desc, Status: model.StatusNew}
	repoOutput := &model.Task{ID: 1, Description: desc, Status: model.StatusNew}
	repo.On("Add", ctx, repoInput).Return(repoOutput, nil)
	task, err := service.Add(ctx, task)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, model.StatusNew, task.Status)
	repo.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	desc := "some test task"
	ID := int64(1)
	repoOutput := &model.Task{ID: ID, Description: desc, Status: model.StatusNew}
	repo.On("Get", ctx, ID).Return(repoOutput, nil)
	task, err := service.Get(ctx, ID)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	repo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	desc := "some test task"
	task := &model.Task{ID: 1, Description: desc, Status: model.StatusInProgress}
	repoInput := &model.Task{ID: 1, Description: desc, Status: model.StatusInProgress}
	repoOutput := &model.Task{ID: 1, Description: desc, Status: model.StatusInProgress}
	repo.On("Update", ctx, repoInput).Return(repoOutput, nil)
	task, err := service.Update(ctx, task)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	repo.AssertExpectations(t)
}

func TestList(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	desc := "some test task"
	repoOutput := []*model.Task{
		&model.Task{ID: 1, Description: desc, Status: model.StatusInProgress},
	}
	repo.On("List", ctx).Return(repoOutput, nil)
	tasks, err := service.List(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	repo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	desc := "some test task"
	task := &model.Task{ID: 1, Description: desc, Status: model.StatusInProgress}
	repo.On("Delete", ctx, task.ID).Return(nil)
	err := service.Delete(ctx, task)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestAssign(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	desc := "some test task"
	task := &model.Task{
		ID: 1, Description: desc,
		Status:   model.StatusInProgress,
		Assigned: []int64{},
	}
	user := &userModel.User{ID: 1}
	users := []*userModel.User{user}
	repoInput := &model.Task{
		ID: 1, Description: desc,
		Status:   model.StatusInProgress,
		Assigned: []int64{user.ID},
	}
	repo.On("Update", ctx, repoInput).Return(repoInput, nil)
	task, err := service.Assign(ctx, task, users...)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	repo.AssertExpectations(t)
}
