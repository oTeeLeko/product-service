package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/product-service/api/controllers"
)

func ProductRoutes(router gin.IRouter, c *controllers.ProductController) {
	group := router.Group("/product")
	group.POST("/", c.CreateProduct)
	group.PATCH("/:id", c.UpdateProduct)
}
