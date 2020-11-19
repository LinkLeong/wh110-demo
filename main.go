package main

import (
	"fmt"
	"net/http"
	"wh110api/pak/logging"
	"wh110api/pak/mgo"
	"wh110api/pak/setting"
	"wh110api/router"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @BasePath /
func main() {
	//model.Setup()
	fmt.Println("开始运行")
	setting.Setup()
	logging.Setup()
	mgo.Init()
	//gredis.Setup()
	r := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServeSetting.HttpPort),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
