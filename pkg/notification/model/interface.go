package model

import (
	"context"
)

// Service - main user service interface
type Service interface {
	SendEmail(ctx context.Context, email, msg string) error
}

//go:generate mockery -name=Service -inpkg
