package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.comElcioAngelo/TrabalhoTCC.git/model"
	"github.comElcioAngelo/TrabalhoTCC.git/usecase"
)

type ProductController struct {
	usecase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController {
		usecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context){

	products,err := p.usecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch products",
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, products)
} 

func (p *ProductController) CreateProduct(ctx *gin.Context){

	var product model.Product

	err := ctx.ShouldBind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while creating product",
			"error": err,
		})
	}
	
}