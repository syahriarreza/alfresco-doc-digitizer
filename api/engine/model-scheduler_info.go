package engine

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// SchedulerInfo SchedulerInfo
// NOTE: there is always one record only in this table
type SchedulerInfo struct {
	Base
	IsRunning bool
}

// BeforeCreate BeforeCreate
func (m *SchedulerInfo) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = "1"
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	}

	return nil
}

// GetIsRunning get is_running status
func (m *SchedulerInfo) GetIsRunning(db *gorm.DB) (bool, error) {
	si := SchedulerInfo{}
	if res := db.First(&si, "1"); res.Error != nil {
		return false, fmt.Errorf("error get scheduler info : %s", res.Error.Error())
	}

	return si.IsRunning, nil
}

// UpdateIsRunning update status is_running
func (m *SchedulerInfo) UpdateIsRunning(db *gorm.DB, isRunning bool) error {
	if res := db.Model(m).Where("id = ?", "1").Update("is_running", isRunning); res.Error != nil {
		//TODO: logging for error
		return fmt.Errorf("error update scheduler is_running : %s", res.Error.Error())
	}

	return nil
}
