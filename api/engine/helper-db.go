package engine

import (
	"errors"
	"fmt"
	"time"

	"github.com/eaciit/toolkit"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB DB
	DB *gorm.DB
)

// InitDB init DB
func InitDB() error {
	// TODO: add mechanism if postgresql is used
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
	)
	fmt.Println("[INIT] DB:", dsn)
	db, e := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		return e
	}
	DB = db

	// init tables
	e = db.AutoMigrate(
		Users{},
		Qrcodes{},
		Jobs{},
		Pages{},
		Histories{},
		SchedulerInfo{},
	)

	if e != nil {
		return e
	}

	// init scheduler_info
	scInfo := SchedulerInfo{}
	if IsNotFound(db.First(&scInfo, "1")) {
		db.Create(&SchedulerInfo{})
	}

	if CountRecord(&Qrcodes{}) == 0 {
		qr := new(Qrcodes)
		qr.ID = "001"
		qr.AlfrescoPath = "/001/001"
		db.Create(qr)
		qr = new(Qrcodes)
		qr.ID = "002"
		qr.AlfrescoPath = "/002/002"
		db.Create(qr)
	}

	return nil
}

// MakeID MakeID
func MakeID(prefix string, l int) string {
	// if l < 20 {
	// 	l = 20
	// }
	p1 := toolkit.Date2String(time.Now(), "YYMMddHHmmss")
	lp2 := l - len(prefix) - len(p1)
	if lp2 <= 0 {
		if prefix == "" {
			return p1
		}
		return fmt.Sprintf("%s%s", prefix, p1)
	}
	p2 := toolkit.GenerateRandomString("ABCDEFGHIJKLMNOPQRSTUVWXYZ", lp2)
	if prefix == "" {
		return fmt.Sprintf("%s%s", p1, p2)
	}
	return fmt.Sprintf("%s%s%s", prefix, p1, p2)
}

// IsNotFound IsNotFound
func IsNotFound(res *gorm.DB) bool {
	return res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound)
}

// CountRecord count based on id [SELECT COUNT(id)]
func CountRecord(modelPtr interface{}, where ...string) int {
	to := struct {
		Total int `json:"total"`
	}{}

	tx := DB.Model(modelPtr).Select("COUNT(id) AS total")

	if len(where) >= 2 {
		whereQry, whereArgs := where[0], where[1:]
		tx = tx.Where(whereQry, toolkit.ToInterfaceArray(whereArgs)...)
	}

	tx = tx.Find(&to)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return 0
	}

	return to.Total
}
