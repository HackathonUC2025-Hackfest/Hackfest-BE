package user

import (
	"time"

	"github.com/google/uuid"
)

type Table struct {
	ID           uuid.UUID    `db:"id"`
	FullName     string       `db:"full_name"`
	Email        string       `db:"email"`
	Password     string       `db:"password"`
	AuthProvider AuthProvider `db:"auth_provider"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
}
