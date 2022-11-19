package entities

import "time"

type Excercises struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
