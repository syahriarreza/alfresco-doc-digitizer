package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"github.com/syahriarreza/alfresco-doc-digitizer/api/engine"
)

func main() {
	if e := readConfig(); e != nil {
		fmt.Println("ERROR:", e)
		return
	}

	// inits
	if e := initFolder(); e != nil {
		fmt.Println("ERROR:", e)
		return
	}

	if e := engine.InitDB(); e != nil {
		fmt.Println("ERROR:", e)
		return
	}

	sch := gocron.NewScheduler(time.UTC)
	se := engine.NewSchedulerEngine(sch)

	if e := se.InitWorkerPool(); e != nil {
		fmt.Println("ERR:", e)
		return
	}

	sch.Every(viper.GetUint64("scheduler.every_x_minute")).Minute().
		Do(se.ScanDropbox)
	sch.StartAsync()

	// router
	router := gin.Default()
	ae := engine.AccountEngine{}
	router.GET("/account", ae.Register)
	router.GET("/account/login", ae.Login)
	router.GET("/extract", ae.Extract)

	if e := router.Run(fmt.Sprintf("%s:%s", "", viper.GetString("application.port"))); e != nil {
		fmt.Printf("ERROR: application is stopping. %s\n", e.Error())
	}
}

func readConfig() error {
	wd, _ := os.Getwd()
	configName := "app"
	fmt.Println("Config:", filepath.Join(wd, "config", configName))

	viper.SetConfigName(configName)
	viper.AddConfigPath(filepath.Join(wd, "config"))
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func initFolder() error {
	folders := viper.GetStringMapString("folder")

	for k, folder := range folders {
		if folder == "" {
			return fmt.Errorf("config 'folder.%s' cannot be empty", k)
		}
		if e := os.MkdirAll(folder, os.ModePerm); e != nil {
			return e
		}
	}

	return nil
}
