package model

import (
	"context"
)

// Service - main user service interface
type Service interface {
	Get(ctx context.Context, id int64) (*User, error)
	Add(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	List(context.Context) ([]*User, error)
	Delete(context.Context, *User) error
}

//go:generate mockery -name=Service -inpkg

// Repository - users repository interface
type Repository interface {
	Get(ctx context.Context, id int64) (*User, error)
	Add(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	List(context.Context) ([]*User, error)
	Delete(ctx context.Context, id int64) error
}

//go:generate mockery -name=Repository -inpkg
