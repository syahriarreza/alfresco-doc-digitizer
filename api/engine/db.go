package engine

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/syahriarreza/alfresco-doc-digitizer/api/helper"
	"github.com/syahriarreza/alfresco-doc-digitizer/api/model"
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
	db, e := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		return e
	}
	DB = db

	// init tables
	e = db.AutoMigrate(
		model.Users{},
		model.Qrcodes{},
		model.Jobs{},
		model.Pages{},
		model.Histories{},
		model.SchedulerInfo{},
	)

	if e != nil {
		return e
	}

	// init scheduler_info
	scInfo := model.SchedulerInfo{}
	if helper.IsNotFound(db.First(&scInfo, "1")) {
		db.Create(&model.SchedulerInfo{})
	}

	return nil
}
