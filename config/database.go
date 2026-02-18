package config

import (
	"log"

	tbl "github.com/oTeeLeko/product-service/models"
	"github.com/oTeeLeko/product-service/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load config: ", err)
	}

	database, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = database.AutoMigrate(&tbl.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	DB = database
}
