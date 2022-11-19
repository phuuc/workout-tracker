package entities

import "time"

type Workout struct {
	ID          int
	UserID      int
	Reps        int
	Sets        int
	ExcerciseID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
