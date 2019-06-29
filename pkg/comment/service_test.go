package comment

import (
	"context"
	"tasks/pkg/comment/model"
	taskModel "tasks/pkg/task/model"
	"tasks/pkg/user"
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
	u := userModel.User{ID: 1, Name: "test user"}
	ctx = user.NewContext(ctx, u)
	task := &taskModel.Task{ID: 2}
	msg := "test comment"
	comment := &model.Comment{ID: 1, Task: task.ID, User: u.ID, Msg: msg}
	repo.On("Add", ctx, task.ID, u.ID, msg).Return(comment, nil)
	comment, err := service.Add(ctx, task, msg)
	assert.NoError(t, err)
	assert.NotNil(t, comment)
	repo.AssertExpectations(t)
}

func TestAddUnAuthorized(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	task := &taskModel.Task{ID: 2}
	msg := "test comment"
	comment, err := service.Add(ctx, task, msg)
	assert.Equal(t, userModel.ErrUnAuthorized, err)
	assert.Nil(t, comment)
	repo.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	u := userModel.User{ID: 1, Name: "test user"}
	task := &taskModel.Task{ID: 2}
	msg := "test comment"
	comment := &model.Comment{ID: 1, Task: task.ID, User: u.ID, Msg: msg}
	repo.On("Get", ctx, comment.ID).Return(comment, nil)
	comment, err := service.Get(ctx, comment.ID)
	assert.NoError(t, err)
	assert.NotNil(t, comment)
	repo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	u := userModel.User{ID: 1, Name: "test user"}
	task := &taskModel.Task{ID: 2}
	msg := "test comment"
	comment := &model.Comment{ID: 1, Task: task.ID, User: u.ID, Msg: msg}
	repo.On("Update", ctx, comment).Return(comment, nil)
	comment, err := service.Update(ctx, comment)
	assert.NoError(t, err)
	assert.NotNil(t, comment)
	repo.AssertExpectations(t)
}

func TestList(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	repoOutput := []*model.Comment{&model.Comment{ID: 1}}
	repo.On("List", ctx).Return(repoOutput, nil)
	tasks, err := service.List(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	repo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	u := userModel.User{Name: "test", Role: userModel.RoleAdmin}
	ctx = user.NewContext(ctx, u)
	comment := &model.Comment{ID: 1}
	repo.On("Delete", ctx, comment.ID).Return(nil)
	err := service.Delete(ctx, comment)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteUnAuthorized(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	comment := &model.Comment{ID: 1}
	err := service.Delete(ctx, comment)
	assert.Equal(t, userModel.ErrUnAuthorized, err)
	repo.AssertExpectations(t)
}

func TestDeletePermissionDenied(t *testing.T) {
	repo, service := getServices()
	ctx := context.Background()
	u := userModel.User{Name: "test", Role: userModel.RoleUser}
	ctx = user.NewContext(ctx, u)
	comment := &model.Comment{ID: 1}
	err := service.Delete(ctx, comment)
	assert.Equal(t, userModel.ErrPermissionDenied, err)
	repo.AssertExpectations(t)
}
