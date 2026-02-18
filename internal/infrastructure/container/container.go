package container

import (
	"github.com/oTeeLeko/product-service/internal/adapter/handler"
	"github.com/oTeeLeko/product-service/internal/adapter/repository/postgres"
	"github.com/oTeeLeko/product-service/internal/usecase"
	"gorm.io/gorm"
)

type Container struct {
	ProductHandler *handler.ProductHandler
}

func NewContainer(db *gorm.DB) *Container {
	// 1. Repository Implementation (Adapter)
	productRepo := postgres.NewProductRepository(db)

	// 2. UseCase (Domain Logic - Constructor Injection)
	productUseCase := usecase.NewProductUseCase(productRepo)

	// 3. Handler (Adapter - Constructor Injection)
	productHandler := handler.NewProductHandler(productUseCase)

	return &Container{
		ProductHandler: productHandler,
	}
}
