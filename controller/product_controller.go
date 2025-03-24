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

	ProductError := p.usecase.CreateProduct(product)
	if ProductError != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error": ProductError.Error(),
		})
		return
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"message": "User created successfully!",
			"user_name": product.Name,
		})
		return 
	}
}

func (p *ProductController) EditProduct(ctx *gin.Context) {
	
	var product model.Product
	var id = product.ID
	var value string 
	
	switch(value) {
	case "" :
	}


	productError := ctx.ShouldBind(id); if productError != nil {
		ctx.JSON(500,gin.H{
			"message": "error while editing product",
			"error": productError,
		})
	}

	ctx.JSON(200,gin.H{
		"message": "product edited sucessfully",
	})
}