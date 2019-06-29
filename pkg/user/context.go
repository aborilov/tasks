package user

import (
	"context"
	"tasks/pkg/user/model"
)

// The contextKey type is unexported to prevent collisions with context keys defined in
// other packages.
type contextKey int

// userKey is the context key for the User. Its value of zero is arbitrary. If
// this package defined other context keys, they would have different integer
// values.
const userKey contextKey = 0

// FromContext extracts the user from ctx, if present.
func FromContext(ctx context.Context) (user model.User, ok bool) {
	// ctx.Value returns nil if ctx has no value for the key; the User type
	// assertion returns ok=false for nil.
	user, ok = ctx.Value(userKey).(model.User)
	return
}

// NewContext returns a new Context carrying user.
func NewContext(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}
