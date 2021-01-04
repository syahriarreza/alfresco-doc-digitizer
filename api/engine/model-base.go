package engine

import (
	"time"
)

// Base model
type Base struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
