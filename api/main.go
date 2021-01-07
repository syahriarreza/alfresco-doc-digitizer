package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/viper"
	"github.com/syahriarreza/alfresco-doc-digitizer/api/engine"
)

func main() {
	if e := readConfig(); e != nil {
		fmt.Println("ERROR:", e)
		return
	}

	// inits
	if e := engine.InitDB(); e != nil {
		fmt.Println("ERROR:", e)
		return
	}

	if e := new(engine.SchedulerEngine).InitWorkerPool(); e != nil {
		fmt.Println("ERR:", e)
		return
	}

	gocron.Every(viper.GetUint64("scheduler.every_x_minute")).Minute().
		Do(new(engine.SchedulerEngine).ScanDropbox)

	// go toolkit.RunCommand("go", "run", filepath.Join("scheduler", "myvent-scheduler.go"))

	// router
	router := gin.Default()
	ae := engine.AccountEngine{}
	router.GET("/account", ae.Register)
	router.GET("/account/login", ae.Login)
	router.GET("/extract", ae.Extract)
	router.Run(":8080")

	fmt.Println("### WILL START CRON")
	<-gocron.Start()
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
