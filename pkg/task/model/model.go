package model

// Task - service layer task model
type Task struct {
	ID          int64
	Status      Status
	Description string
	Assigned    []int64
}

func (t *Task) String() string {
	return t.Description
}
