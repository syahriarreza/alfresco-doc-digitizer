package model

import (
	"time"

	"github.com/syahriarreza/alfresco-doc-digitizer/api/helper"
	"gorm.io/gorm"
)

// Pages Pages
type Pages struct {
	Base
	JobID    string
	Filename string
}

// BeforeCreate BeforeCreate
func (m *Pages) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = helper.MakeID("", 16)
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	}

	return nil
}
