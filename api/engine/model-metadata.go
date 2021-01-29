package engine

import (
	"time"

	"gorm.io/gorm"
)

// Metadata Metadata
type Metadata struct {
	Base
	QrcodeID string
	JobID    string
	Field    string
	Value    string
}

// BeforeCreate BeforeCreate
func (m *Metadata) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = MakeID("", 16)
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	}

	return nil
}
