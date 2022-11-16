package entities

import "time"

type Workout struct {
	ID        int
	UserID    int
	Excerise  []Excerise
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Excerise struct {
	Reps    int
	Sets    int
	TypesID int
}
