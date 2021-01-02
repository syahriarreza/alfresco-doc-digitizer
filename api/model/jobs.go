package model

import (
	"time"

	"github.com/syahriarreza/alfresco-doc-digitizer/api/helper"
	"gorm.io/gorm"
)

const (
	// JobInQueue status
	JobInQueue string = "in_queue" // pdf sedang / akan di scan qrcode nya
	// JobScanning status
	JobScanning string = "scanning" // pdf sedang / akan di scan qrcode nya
	// JobInReview status
	JobInReview string = "in_review" // sedang di review dari user, sbelum di upload
	// JobUploading status
	JobUploading string = "uploading" // sedang process upload ke alfresco
	// JobUploaded status
	JobUploaded string = "uploaded" // berhasil upload ke alfresco
)

// JobStatuses status list
var JobStatuses = map[string]string{
	JobInQueue:   "In Queue",
	JobScanning:  "Scanning",
	JobInReview:  "In Review",
	JobUploading: "Uplading",
	JobUploaded:  "Uploaded",
}

// Jobs Jobs
type Jobs struct {
	Base
	QrcodeID         string
	AlfrescoFilename string // param pdf filename utk webscript uploader
	PdfFilename      string // actual pdf filename
	Status           string // job's status
}

// BeforeCreate hook
func (m *Jobs) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = helper.MakeID("", 16)
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	}

	return nil
}

// Run run job
func (m *Jobs) Run() error {

	return nil
}

func (m *Jobs) runScan() error {

	return nil
}
