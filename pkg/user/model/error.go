package model

import "errors"

// ErrPermissionDenied - raise when user does not have required permissions
var ErrPermissionDenied = errors.New("permission denied")

// ErrUnAuthorized - raise when user does not authorized
var ErrUnAuthorized = errors.New("unauthorized")
