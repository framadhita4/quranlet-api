package database

import (
	"fmt"
	"strconv"

	"github.com/framadhita4/quranlet-api/app/models"
	"github.com/framadhita4/quranlet-api/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := configs.GetEnv("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("DB_PORT must be number")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configs.GetEnv("DB_HOST"), port, configs.GetEnv("DB_USER"), configs.GetEnv("DB_PASSWORD"), configs.GetEnv("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Surah{})

	fmt.Println("Connection Opened to Database")
}
