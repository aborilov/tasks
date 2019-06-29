package model

// Status - task statuses
type Status uint8

//go:generate enumer -type=Status -text -json -trimprefix=Status -transform=snake -output=enum_status_gen.go

const (
	// StatusNew - task just created
	StatusNew Status = 1 + iota
	// StatusInProgress - work on task is in progress
	StatusInProgress
	// StatusCompleted - task completed
	StatusCompleted
	// StatusArchived - task archived
	StatusArchived
)
