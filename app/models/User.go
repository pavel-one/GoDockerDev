package models

import (
	"time"
)

type User struct {
	ID        uint
	Username  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
