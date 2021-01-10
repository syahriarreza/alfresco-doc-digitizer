package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-co-op/gocron"
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
	schedule *gocron.Scheduler
}

// NewSchedulerEngine create scheduler engine instance
func NewSchedulerEngine(schedule *gocron.Scheduler) *SchedulerEngine {
	se := new(SchedulerEngine)
	se.schedule = schedule
	return se
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
func (eg *SchedulerEngine) ScanDropbox() {
	if e := eg.scanDropBox(); e != nil {
		fmt.Printf("[ERROR] %s\n", e.Error())
	}
}

func (eg *SchedulerEngine) scanDropBox() error {
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
	for _, f := range fs {
		//TODO: checking for active file

		if isPDF, _ := filepath.Match("*.pdf", f.Name()); isPDF && !f.IsDir() {
			newJobID := MakeID("", 16)
			newFilename := fmt.Sprintf("SC-%s.pdf", newJobID)

			if e := os.Rename(
				filepath.Join(viper.GetString("folder.dropbox"), f.Name()),
				filepath.Join(viper.GetString("folder.scan"), newFilename),
			); e != nil {
				return fmt.Errorf("error move and rename file '%s' : %s", f.Name(), e.Error())
			}

			newJob := Jobs{
				PdfFilename: newFilename,
				Status:      JobPending,
				IsInQueue:   false,
			}
			newJob.ID = newJobID
			DB.Create(&newJob)
		}
	}

	// enqueue pending jobs (jobs thats just created & existing jobs which has status IsInQueue: false)
	pendingJobs := []Jobs{}
	if res := DB.Where("is_in_queue = ? AND status != ?", false, JobInReview).Find(&pendingJobs); res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("error get pending jobs : %s", res.Error.Error())
	}

	for _, job := range pendingJobs {
		job.IsInQueue = true
		if res := DB.Save(&job); res.Error != nil {
			return res.Error
		}
		if e := Pool.Invoke(&job); e != nil {
			return e
		}
	}

	_, time := eg.schedule.NextRun()
	fmt.Println("CRON | Finished, Next Run:", time)
	//TODO: next run nya ga bener..
	return nil
}
