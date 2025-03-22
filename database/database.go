package database

import (
	"fmt"
	"log"

	"iot_api_with_go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/iot_pertemuan_3?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	fmt.Println("Database terkoneksi")

	// Auto migrate tabel
	DB.AutoMigrate(&models.SensorValue{}, &models.SensorData{})
}

func GetDB() *gorm.DB {
	return DB
}
