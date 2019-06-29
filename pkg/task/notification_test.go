package task

import (
	"context"
	notificationModel "tasks/pkg/notification/model"
	"tasks/pkg/task/model"
	userModel "tasks/pkg/user/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var anyString = mock.AnythingOfType("string")

func getNotificationServices() (model.Service, *model.MockService, *userModel.MockService, *notificationModel.MockService) {
	service := &model.MockService{}
	notification := &notificationModel.MockService{}
	user := &userModel.MockService{}
	taskNotification := NewNotificationService(service, notification, user)
	return taskNotification, service, user, notification
}

func TestNotificationAdd(t *testing.T) {
	notificationService, service, user, notification := getNotificationServices()
	ctx := context.Background()
	desc := "some test task"
	task := &model.Task{Description: desc}
	service.On("Add", ctx, task).Return(task, nil)
	task, err := notificationService.Add(ctx, task)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	service.AssertExpectations(t)
	user.AssertExpectations(t)
	notification.AssertExpectations(t)
}

func TestNotificationUpdate(t *testing.T) {
	notificationService, service, user, notification := getNotificationServices()
	ctx := context.Background()
	desc := "some test task"
	u := &userModel.User{ID: 1, Name: "test", Email: "test@test.com"}
	task := &model.Task{Description: desc, Assigned: []int64{u.ID}}
	service.On("Update", ctx, task).Return(task, nil)
	user.On("Get", ctx, u.ID).Return(u, nil)
	notification.On("SendEmail", ctx, u.Email, anyString).Return(nil)
	task, err := notificationService.Update(ctx, task)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	service.AssertExpectations(t)
	user.AssertExpectations(t)
	notification.AssertExpectations(t)
}

func TestNotificationGet(t *testing.T) {
	notificationService, service, user, notification := getNotificationServices()
	ctx := context.Background()
	id := int64(1)
	desc := "some test task"
	task := &model.Task{Description: desc}
	service.On("Get", ctx, id).Return(task, nil)
	task, err := notificationService.Get(ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	service.AssertExpectations(t)
	user.AssertExpectations(t)
	notification.AssertExpectations(t)
}

func TestNotificationList(t *testing.T) {
	notificationService, service, user, notification := getNotificationServices()
	ctx := context.Background()
	service.On("List", ctx).Return([]*model.Task{}, nil)
	task, err := notificationService.List(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	service.AssertExpectations(t)
	user.AssertExpectations(t)
	notification.AssertExpectations(t)
}

func TestNotificationDelete(t *testing.T) {
	notificationService, service, user, notification := getNotificationServices()
	ctx := context.Background()
	task := &model.Task{ID: 1}
	service.On("Delete", ctx, task).Return(nil)
	err := notificationService.Delete(ctx, task)
	assert.NoError(t, err)
	service.AssertExpectations(t)
	user.AssertExpectations(t)
	notification.AssertExpectations(t)
}

func TestNotificationAssign(t *testing.T) {
	notificationService, service, user, notification := getNotificationServices()
	ctx := context.Background()
	task := &model.Task{ID: 1}
	u := &userModel.User{ID: 1}
	service.On("Assign", ctx, task, u).Return(task, nil)
	task, err := notificationService.Assign(ctx, task, u)
	assert.NoError(t, err)
	service.AssertExpectations(t)
	user.AssertExpectations(t)
	notification.AssertExpectations(t)
}
