package controller

import (
	"net/http"
	"strconv"

	"goapi/model"
	"goapi/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductController(usecase *usecase.ProductUsecase) *ProductController {
	return &ProductController{
		productUsecase: usecase,
	}
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	products, err := c.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetProductByID(ctx *gin.Context) {
	queryID := ctx.Query("id")
	if queryID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	productID, _ := strconv.Atoi(queryID)

	product, err := c.productUsecase.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) InsertProduct(ctx *gin.Context) {
	var product model.Product

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := c.productUsecase.InsertProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, insertedProduct)
}
