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


	ctx.JSON(http.StatusOK,gin.H{
		"message": "product created successfully",
	})

}

func (p *ProductController) EditProduct(ctx *gin.Context) {
	id := ctx.Param("product_id")

	var product model.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}


	// * Alterando o product de acordo com os campos retornados.
	var fields []string
	var values []interface{}

	if product.Name != "" {
		fields = append(fields, "name = ?")
		values = append(values, product.Name)
	}
	if product.Price != 0.0 {
		fields = append(fields, "name = ?")
		values = append(values, product.Price)
	}
	if product.Description != "" {
		fields = append(fields, "name = ?")
		values = append(values, product.Description)
	}
	if product.CategoryID != 0 {
		fields = append(fields, "name = ?")
		values = append(values, product.CategoryID)
	}
	if product.BrandID != 0 {
		fields = append(fields, "name = ?")
		values = append(values, product.BrandID)
	}

	if len(fields) == 0 {
		ctx.JSON(400, gin.H{
			"error": "no fields to update",
		})
	}

	productError := p.usecase.EditProduct(fields,values,id)
	if productError != nil {
		ctx.JSON(500,gin.H{
			"message": "error while editing product",
			"error": productError,
		})
	}

	ctx.JSON(200,gin.H{
		"message": "product edited sucessfully",
	})
}
