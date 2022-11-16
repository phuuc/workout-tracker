package entities

import "time"

type Users struct {
	ID           int
	Email        string
	HashedPasswd []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
