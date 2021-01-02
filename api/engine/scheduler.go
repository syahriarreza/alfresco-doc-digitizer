package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/panjf2000/ants"
	"github.com/spf13/viper"
	"github.com/syahriarreza/alfresco-doc-digitizer/api/helper"
	"github.com/syahriarreza/alfresco-doc-digitizer/api/model"
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
			if v, ok := payload.(*model.Jobs); ok {
				// do something
				fmt.Println(v)
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
	scInfo := new(model.SchedulerInfo)

	//check is_running
	isRunning, e := scInfo.GetIsRunning(DB)
	if isRunning {
		fmt.Println("scheduler is currently running") //TODO: remove this
		return nil
	}
	if e != nil {
		//TODO: logging for error
		return e
	}

	fs, e := ioutil.ReadDir(viper.GetString("folder.dropbox"))
	if e != nil {
		return fmt.Errorf("error reading folder dropbox : %s", e.Error())
	}

	if e := scInfo.UpdateIsRunning(DB, true); e != nil {
		//TODO: logging for error
		return e
	}

	// scan all files
	for i, f := range fs {
		//TODO: checking for active file

		if isPDF, _ := filepath.Match("*.pdf", f.Name()); isPDF && !f.IsDir() {
			fmt.Println(fmt.Sprintf("File #%d: %s", i+1, f.Name()))

			// move & rename pdf
			newFilename := helper.MakeID("", 12) + ".pdf"
			if e := os.Rename(
				filepath.Join(viper.GetString("folder.dropbox"), f.Name()),
				filepath.Join(viper.GetString("folder.scan"), newFilename),
			); e != nil {
				return fmt.Errorf("error move and rename file '%s' : %s", f.Name(), e.Error())
			}

			DB.Create(&model.Jobs{
				PdfFilename: newFilename,
				Status:      model.JobInQueue,
			})
		}
	}

	if e := scInfo.UpdateIsRunning(DB, false); e != nil {
		//TODO: logging for error
		return e
	}

	//TODO: get all jobs with status == scan
	//TODO: enqueue jobs

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
