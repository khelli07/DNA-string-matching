package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (db *gorm.DB) {
	password := "" // GANTI PASSWORD DB
	dsn := "root:" + password + "@tcp(127.0.0.1:3306)/gosql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
