package entity

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Age       int
	IsStudent bool
	CreatedAt time.Time
}
