package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/product-service/internal/infrastructure/container"
	"github.com/oTeeLeko/product-service/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(c *container.Container) *gin.Engine {
	r := gin.Default()

	// Swagger
	r.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	// Middleware
	r.Use(middleware.AccessLogger())
	r.Use(middleware.ErrorLogger())

	// Product Routes
	productRoutes := r.Group("/product")
	{
		productRoutes.POST("", c.ProductHandler.CreateProduct)
		productRoutes.PATCH("/:id", c.ProductHandler.UpdateProduct)
	}

	return r
}
