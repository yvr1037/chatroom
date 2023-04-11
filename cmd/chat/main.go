package main

import (
	"chatroom/global"
	"chatroom/internal/routers"
	"chatroom/internal/setting"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func setupSettings() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	}
	err = setting.ReadSection("DataBase", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSettings.ReadTimeout *= time.Second
	global.ServerSettings.WriteTimeout *= time.Second
	return nil
}

func init() {
	err := setupSettings()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSettings.Runmode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSettings.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSettings.ReadTimeout,
		WriteTimeout:   global.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe();err != nil {
		log.Fatalf("ListenAndServer err: %v",err)
	}
}
