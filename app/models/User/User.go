package User

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint
	Name        string
	Email       *string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
