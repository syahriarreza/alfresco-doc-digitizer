package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/panjf2000/ants"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// Pool ants pool
	Pool *ants.PoolWithFunc
)

// SchedulerEngine engine
type SchedulerEngine struct {
}

// InitWorkerPool init ants pool
func (eg *SchedulerEngine) InitWorkerPool() error {
	pool, e := ants.NewPoolWithFunc(viper.GetInt("scheduler.poolsize"), func(payload interface{}) {
		if payload != nil {
			if v, ok := payload.(*Jobs); ok {
				v.Run(DB)
			}
		}
	})
	if e != nil {
		return e
	}

	Pool = pool

	return nil
}

// ScanDropbox regular func for cron
func (eg *SchedulerEngine) ScanDropbox() error {
	scInfo := new(SchedulerInfo)

	//check is_running
	isRunning, e := scInfo.GetIsRunning(DB)
	if isRunning {
		fmt.Println("scheduler is still running") //TODO: remove this
		return nil
	}
	if e != nil {
		//TODO: logging for error
		return e
	}

	if e := scInfo.UpdateIsRunning(DB, true); e != nil {
		return e
	}
	defer scInfo.UpdateIsRunning(DB, false)

	fs, e := ioutil.ReadDir(viper.GetString("folder.dropbox"))
	if e != nil {
		return fmt.Errorf("error reading folder dropbox : %s", e.Error())
	}

	// scan all files
	for i, f := range fs {
		//TODO: checking for active file
		if isPDF, _ := filepath.Match("*.pdf", f.Name()); isPDF && !f.IsDir() {
			fmt.Println(fmt.Sprintf("File #%d: %s", i+1, f.Name()))

			// move & rename pdf
			newFilename := MakeID("", 12) + ".pdf"
			if e := os.Rename(
				filepath.Join(viper.GetString("folder.dropbox"), f.Name()),
				filepath.Join(viper.GetString("folder.scan"), newFilename),
			); e != nil {
				return fmt.Errorf("error move and rename file '%s' : %s", f.Name(), e.Error())
			}

			DB.Create(&Jobs{
				PdfFilename: newFilename,
				Status:      JobPending,
				IsInQueue:   false,
			})
		}
	}

	// enqueue pending jobs
	pendingJobs := []Jobs{}
	if res := DB.Where("is_in_queue = ? AND status != ?", false, JobInReview).Find(&pendingJobs); res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("error get pending jobs : %s", res.Error.Error())
	}

	for _, job := range pendingJobs {
		job.IsInQueue = true
		if res := DB.Save(&job); res.Error != nil {
			return res.Error
		}
		if e := Pool.Invoke(job); e != nil {
			return e
		}
	}

	return nil
}

// ReadQRCodeInPDF ReadQRCodeInPDF
func (eg *SchedulerEngine) ReadQRCodeInPDF() error {
	qrDecoder := qrcode.NewQRCodeReader()
	doc, e := fitz.New(viper.GetString("folder.dropbox"))
	if e != nil {
		panic(e)
	}
	defer doc.Close()

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, e := doc.Image(n)
		if e != nil {
			return fmt.Errorf("get image: %s", e.Error())
		}

		bmp, e := gozxing.NewBinaryBitmapFromImage(img)
		if e != nil {
			return fmt.Errorf("create new binary bitmap: %s", e.Error())
		}

		res, _ := qrDecoder.Decode(bmp, nil)
		if res == nil {
			fmt.Printf("Page #%03d: QR Code not found\n", n+1)
		} else {
			fmt.Printf("Page #%03d: %s\n", n+1, res.String())
		}
	}

	return nil
}
