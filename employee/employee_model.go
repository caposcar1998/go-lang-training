package employee

import (
	"time"
)

type Position int

const (
	Undetermined Position = iota
	Junior
	Senior
	Manager
	CEO
)

type Employee struct {
	ID          int
	FullName    string
	Position    Position
	Salary      float64
	Joined      time.Time
	OnProbation bool
	CreatedAt   time.Time
}
