package model

import (
	"time"

	"github.com/syahriarreza/alfresco-doc-digitizer/api/helper"
	"gorm.io/gorm"
)

// Users Users
type Users struct {
	Base
	Name     string
	Username string
	Password string
}

// BeforeCreate BeforeCreate
func (m *Users) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = helper.MakeID("", 16)
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	}

	return nil
}
