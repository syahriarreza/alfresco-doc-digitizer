package engine

import (
	"time"

	"gorm.io/gorm"
)

const (
	// JobPending status
	JobPending string = "pending" // job belum masuk queue
	// JobScanning status
	JobScanning string = "scanning" // pdf sedang di scan qrcode nya
	// JobInReview status
	JobInReview string = "in_review" // sedang di review dari user, sbelum di upload
	// JobUploading status
	JobUploading string = "uploading" // sedang process upload ke alfresco
)

// JobStatuses status list
var JobStatuses = map[string]string{
	JobPending:   "Pending",
	JobScanning:  "Scanning",
	JobInReview:  "In Review",
	JobUploading: "Uploading",
}

// Jobs Jobs
type Jobs struct {
	Base
	QrcodeID         string
	AlfrescoFilename string // param pdf filename utk webscript uploader
	PdfFilename      string // actual pdf filename
	Status           string // job's status
	IsInQueue        bool   // is job already in queue
}

// BeforeCreate hook
func (m *Jobs) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = MakeID("", 16)
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	}

	return nil
}

// Run run job
func (m *Jobs) Run(db *gorm.DB) error {
	switch true {
	case m.Status == JobPending || m.Status == JobScanning:
		if e := m.runScan(db); e != nil {
			m.IsInQueue = false
			db.Save(m)
			return e
		}
		if res := db.Delete(m); res.Error != nil {
			return res.Error
		}

	case m.Status == JobUploading:
		if e := m.runUpload(db); e != nil {
			return e
		}
		m.IsInQueue = false
		if res := db.Save(m); res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (m *Jobs) runScan(db *gorm.DB) error {
	m.Status = JobScanning
	if res := db.Save(m); res.Error != nil {
		return res.Error
	}

	pdfExts, e := ExtractPDF(m.PdfFilename)
	if e != nil {
		return e
	}

	// create multiple jobs based on pdf extracts
	for _, pe := range pdfExts {
		nj := new(Jobs)
		nj.QrcodeID = pe.QRCode.ID
		nj.AlfrescoFilename = pe.QRCode.AlfrescoDefaultFilename

		if nj.QrcodeID == "" {
			//TODO: handle file kalo ga ada qr code nya

		} else if pe.QRCode.IsReviewNeeded {
			nj.PdfFilename = ""
			nj.Status = JobInReview
			//TODO: create pages and save image to file

		} else {

			nj.Status = JobUploading
			nj.PdfFilename = MakeID("", 16) + ".pdf"
			//TODO: create pdf file (nj.PdfFilename) in 'upload' folder from images (pe.Pages)
		}

		//TODO: save new job
	}

	return nil
}

func (m *Jobs) runUpload(db *gorm.DB) error {
	m.Status = JobUploading
	if res := db.Save(m); res.Error != nil {
		return res.Error
	}

	//TODO: Upload to alfresco

	//TODO: if success, record nya di delete & + histories

	return nil
}
