package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ziglu/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

// Setup initializes the database instance
func CreateConnection() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable", setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Port, setting.DatabaseSetting.Database)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database connection to %s:%s - %s", setting.DatabaseSetting.Host, setting.DatabaseSetting.Port, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("[INFO] DB connection successful")
}

func CloseConnection() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}
