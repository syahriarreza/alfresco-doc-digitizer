package engine

import (
	"time"

	"gorm.io/gorm"
)

// Qrcodes Qrcodes
type Qrcodes struct {
	Base
	AlfrescoPath            string
	AlfrescoDefaultFilename string // default name ketika di upload ke alfresco
	IsReviewNeeded          bool
}

// BeforeCreate BeforeCreate
func (m *Qrcodes) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = MakeID("", 16)

		if m.CreatedAt.IsZero() {
			m.CreatedAt = time.Now()
			m.UpdatedAt = time.Now()
		}
	}

	return nil
}
