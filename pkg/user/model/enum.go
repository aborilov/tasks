package model

// Role - user role
type Role uint8

//go:generate enumer -type=Role -text -json -trimprefix=Role -transform=snake -output=enum_role_gen.go

const (
	// RoleAdmin - admin user role
	RoleAdmin Role = 1 + iota
	// RoleUser - simple user role
	RoleUser
)
