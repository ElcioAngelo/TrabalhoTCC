package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"trabalhoTcc.com/mod/model"
	"trabalhoTcc.com/mod/repository"
)

type ProductController struct {
	repository repository.ProductRepository 
}

func NewProductController(repository repository.ProductRepository) ProductController {
	return ProductController {
		repository: repository,
	}
}

func (p *ProductController) GetProduct(ctx *gin.Context){
	requestId := ctx.Param("product_id");

	id,err := strconv.Atoi(requestId);
	if err != nil {
		ctx.JSON(http.StatusBadGateway,gin.H{
			"message": "failed to parse int",
			"error": err.Error(),
		})
	}

	product,err := p.repository.GetProductById(id);
	if err != nil {
		ctx.JSON(http.StatusBadGateway,gin.H{
			"message": "failed to parse int",
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,product)
}

func (p *ProductController) GetProducts(ctx *gin.Context){

	products,err := p.repository.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch products",
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductsAdmin(ctx *gin.Context){

	products,err := p.repository.GetProductsAll()
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

	ProductError := p.repository.CreateProduct(product)
	if ProductError != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message": "error while creating product",
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

type value_update struct {
	id 				int		`json:"id"`
	value 			string 	`json:"value"`
	update_type 	string 	`json:"update_type"`
}

// !! ALTERAR NOME DOS PRODUTOS 
func(p * ProductController) EditProductName(ctx *gin.Context) {
	type value struct {
		Name string `json:"name"`
	}

	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}
 
	var request value 

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "error while parsing product name",
			"error": err.Error(),
		})
	}

	producterr := p.repository.EditProductName(id,request.Name)
	if producterr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while altering product name",
			"error": producterr.Error(),
		})
	}

	ctx.JSON(200,gin.H{
		"message": "sucessfully updated product name",
	})
}
// !! ALTERAR PREÇO DOS PRODUTOS
func(p * ProductController) EditProductPrice(ctx *gin.Context) {
	type value struct {
		Price string `json:"price"`
	}

	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}
	var request value

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "error while parsing product price",
			"error": err.Error(),
		})
	}

	producterr := p.repository.EditProductPrice(id,request.Price)
	if producterr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while altering product price",
			"error": producterr.Error(),
		})
	}
	
	ctx.JSON(200,gin.H{
		"message": "sucessfully updated product price",
	})
}

func(p * ProductController) EditProductDescription(ctx *gin.Context) {
	type value struct {
		Description string `json:"description"`
	}

	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}
	var request value

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "error while parsing product description",
			"error": err.Error(),
		})
	}

	producterr := p.repository.EditProductDescription(id,request.Description)
	if producterr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while altering product description",
			"error": producterr.Error(),
		})
	}
	
	ctx.JSON(200,gin.H{
		"message": "sucessfully updated product description",
	})
}

func (p *ProductController) FindProductByCategory(ctx *gin.Context) {
	category := ctx.Param("category");

	products, err := p.repository.SearchProductByCategory(category);
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "unable to fetch products by category.",
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductSales(ctx *gin.Context){

	sales, err := p.repository.GetSales();
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "unable to fetch sales",
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,sales);
}
