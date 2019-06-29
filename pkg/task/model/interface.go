package model

import "context"
import userModel "tasks/pkg/user/model"

// Service - main task service interface
type Service interface {
	Get(ctx context.Context, id int64) (*Task, error)
	Add(context.Context, *Task) (*Task, error)
	Update(context.Context, *Task) (*Task, error)
	List(context.Context) ([]*Task, error)
	Delete(context.Context, *Task) error
	Assign(context.Context, *Task, ...*userModel.User) (*Task, error)
}

//go:generate mockery -name=Service -inpkg

// Repository - tasks repository interface
type Repository interface {
	Get(ctx context.Context, id int64) (*Task, error)
	Add(context.Context, *Task) (*Task, error)
	Update(context.Context, *Task) (*Task, error)
	List(context.Context) ([]*Task, error)
	Delete(ctx context.Context, id int64) error
}

//go:generate mockery -name=Repository -inpkg
