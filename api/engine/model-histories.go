package engine

import (
	"time"

	"gorm.io/gorm"
)

// Histories Histories
type Histories struct {
	Base
	QrcodeID  string
	Filename  string
	IsSuccess bool
	Error     string
}

// BeforeCreate BeforeCreate
func (m *Histories) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = MakeID("", 16)
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	}

	return nil
}
