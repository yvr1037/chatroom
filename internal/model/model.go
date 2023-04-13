package model

import (
	"chatroom/internal/setting"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	dsn := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,   // ip+port
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)  // 最大空闲连接数
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)  // 最大打开连接数
	return db, nil
}
