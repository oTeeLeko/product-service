// @title Product Service API
// @version 1.0
// @description This is a sample product service API
// @BasePath /
package main

import (
	"log"

	"github.com/oTeeLeko/product-service/api"
	"github.com/oTeeLeko/product-service/config"
	"github.com/oTeeLeko/product-service/core/store"
	"github.com/oTeeLeko/product-service/utils"
)

func main() {
	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load config: ", err)
	}

	config.ConnectDatabase()
	db := config.DB

	store := store.NewStore(db)
	server, err := api.NewServer(cfg, store)
	if err != nil {
		log.Fatal("Unable to create server: ", err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
