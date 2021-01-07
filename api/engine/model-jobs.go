package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/viper"
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
	newJobs := []Jobs{}
	for _, pe := range pdfExts {
		nj := new(Jobs)
		nj.Base.ID = MakeID("", 16)
		nj.QrcodeID = pe.QRCode.ID
		nj.AlfrescoFilename = pe.QRCode.AlfrescoDefaultFilename

		if pe.QRCode.IsReviewNeeded {
			// Review
			nj.PdfFilename = ""
			nj.Status = JobInReview
			for pgI, pg := range pe.Pages {
				pageFileName := fmt.Sprintf("%s-%s.pdf", MakeID("", 16), FormatNumberDigit((pgI+1), 4))

				// create Pages & save image
				nPage := new(Pages)
				nPage.JobID = nj.Base.ID
				nPage.Filename = pageFileName
				if res := db.Create(nPage); res.Error != nil {
					return res.Error
				}

				if e := SaveImage(pg, filepath.Join(viper.GetString("folder.process"), pageFileName)); e != nil {
					return e
				}
			}
		} else {
			// Direct Upload
			nj.Status = JobUploading
			nj.PdfFilename = MakeID("", 16) + ".pdf"

			//TODO: set metadata based on QRCode (pe.QRCode)

			if nj.QrcodeID == "" {
				//TODO: set based on app.yml
			}

			// create pdf & save temp image
			pdf := gofpdf.New("P", "mm", "A4", "")
			imgFilePaths := []string{}
			for pgI, pg := range pe.Pages {
				pageFileName := fmt.Sprintf("%s-%s.pdf", MakeID("", 16), FormatNumberDigit((pgI+1), 4))
				imgFilePath := filepath.Join(viper.GetString("folder.process"), pageFileName)
				if e := SaveImage(pg, imgFilePath); e != nil {
					return e
				}
				imgFilePaths = append(imgFilePaths, imgFilePath)

				pdf.AddPage()
				pdf.SetFont("Arial", "B", 16)
				pdf.Image(imgFilePath, 0, 0, 0, 0, false, "", 0, "")
			}

			pdfFilePath := filepath.Join(viper.GetString("folder.upload"), nj.PdfFilename)
			if e := pdf.OutputFileAndClose(pdfFilePath); e != nil {
				return e
			}
			for _, imgfp := range imgFilePaths {
				os.Remove(imgfp)
			}
		}

		newJobs = append(newJobs, *nj)
	}

	if res := db.Create(&newJobs); res.Error != nil {
		return res.Error
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
