package model

import (
	"context"
	taskModel "tasks/pkg/task/model"
)

// Service - main comment service interface
type Service interface {
	Get(ctx context.Context, id int64) (*Comment, error)
	Add(context.Context, *taskModel.Task, string) (*Comment, error)
	Update(context.Context, *Comment) (*Comment, error)
	List(context.Context) ([]*Comment, error)
	Delete(context.Context, *Comment) error
}

//go:generate mockery -name=Service -inpkg

// Repository - users repository interface
type Repository interface {
	Get(ctx context.Context, id int64) (*Comment, error)
	Add(ctx context.Context, taskID, userID int64, msg string) (*Comment, error)
	Update(context.Context, *Comment) (*Comment, error)
	List(context.Context) ([]*Comment, error)
	Delete(ctx context.Context, id int64) error
}

//go:generate mockery -name=Repository -inpkg
