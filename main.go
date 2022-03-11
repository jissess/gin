package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"myGin/common"
	"myGin/router"
	"os"
)

func main() {
	InitConfig()
	db, err := common.InitDB()
	if err != nil {
		panic(err.Error())
		return
	}
	defer db.Close()
	r := gin.Default()
	router.InitRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDit, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDit + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
