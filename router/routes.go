package router

import (
	"goapi/config"
	"goapi/controller"
	"goapi/repository"
	"goapi/usecase"

	"github.com/gin-gonic/gin"
)

const basePath = "/api"

func initRoutes(router *gin.Engine) {
	ProductRepository := repository.NewProductRepository(config.Database)

	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	ProductController := controller.NewProductController(ProductUsecase)

	api := router.Group(basePath)
	{
		api.GET("/products", ProductController.GetProducts)
		api.GET("/product", ProductController.GetProductByID)
		api.POST("/product", ProductController.InsertProduct)
	}
}
