// @title Product Service API
// @version 1.0
// @description This is a sample product service API with Clean Architecture
// @BasePath /
package main

import (
	"log"

	"github.com/oTeeLeko/product-service/config"
	_ "github.com/oTeeLeko/product-service/docs" // Swagger docs
	"github.com/oTeeLeko/product-service/internal/domain/entity"
	"github.com/oTeeLeko/product-service/internal/infrastructure/container"
	"github.com/oTeeLeko/product-service/internal/infrastructure/router"
	"github.com/oTeeLeko/product-service/utils"
)

func main() {
	// 1. Load Config
	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 2. Connect DB
	config.ConnectDatabase()
	db := config.DB

	// 3. Auto Migrate (Optional, but good for project start)
	db.AutoMigrate(&entity.Product{})

	// 4. Initialize DI Container
	diContainer := container.NewContainer(db)

	// 5. Setup Router
	r := router.SetupRouter(diContainer)

	// 6. Start Server
	if err := r.Run(cfg.HTTPServerAddress); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
