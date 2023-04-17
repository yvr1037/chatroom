package main

import (
	"chatroom/global"
	"chatroom/internal/model"
	"chatroom/internal/routers"
	"chatroom/internal/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	err = setting.ReadSection("DataBase", &global.DatabaseSettings)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT",global.JWTSettings)
	if err != nil {
		return err
	}
	global.ServerSettings.ReadTimeout *= time.Second
	global.ServerSettings.WriteTimeout *= time.Second
	global.JWTSettings.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DatabaseSettings)
	if err != nil {
		return err
	}
	return nil
}

func setupTableModel(db *gorm.DB,models ...interface{}) error {
	err := db.AutoMigrate(models)
	if err != nil {
		return err
	}
	return nil 
}

func init() {
	err := setupSettings()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v",err)
	}
	// err = setupTableModel(global.DBEngine,&model.User{},&model.Message{})
	err = setupTableModel(global.DBEngine, &model.User{})
	// err = setupTableModel(global.DBEngine, &model.Message{})
	if err != nil {
		log.Fatalf("init.setupTableModel err: %v",err)
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
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe err: %v", err)
	}
}
