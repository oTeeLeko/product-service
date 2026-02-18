package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/product-service/api/controllers"
	"github.com/oTeeLeko/product-service/api/routes"
	"github.com/oTeeLeko/product-service/core/store"
	_ "github.com/oTeeLeko/product-service/docs"
	"github.com/oTeeLeko/product-service/middleware"
	"github.com/oTeeLeko/product-service/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config utils.Config
	store  *store.Store
	router *gin.Engine
}

func NewServer(config utils.Config, store *store.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	router.Use(middleware.AccessLogger())
	router.Use(middleware.ErrorLogger())

	routes.ProductRoutes(router, controllers.NewProductController(s.store.ProductRepo))
	s.router = router
}

func (s *Server) Start() error {
	return s.router.Run(s.config.HTTPServerAddress)
}
