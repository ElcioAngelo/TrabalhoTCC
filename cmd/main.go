package main

import (
	"github.com/gin-gonic/gin"
	"github.comElcioAngelo/TrabalhoTCC.git/controller"
	"github.comElcioAngelo/TrabalhoTCC.git/db"
	"github.comElcioAngelo/TrabalhoTCC.git/repository"
	"github.comElcioAngelo/TrabalhoTCC.git/usecase"
)



func main() {
	server := gin.Default()
	
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// * Usuários ################################################################################################
	
	UserRepository := repository.NewUserRepository(dbConnection)
	UserUsecase := usecase.NewUserRepository(UserRepository)
	UserController := controller.NewUserController(UserUsecase)


	// * Produtos ################################################################################################
	
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	// * Requisições GET ################################################################################################
	
	server.GET("/fetch/user/:user_id",UserController.GetUser)
	server.GET("/fetch/products", ProductController.GetProducts)
	
	// * Requisições POST ################################################################################################
	
	server.POST("/create/user", UserController.CreateUser)
	server.POST("/create/product",ProductController.CreateProduct)
	
	// * Requisições PATCH ################################################################################################
	
	server.POST("/edit/Product/name",ProductController.EditProduct)
	
	// * Requisições DELETE ################################################################################################

	server.DELETE("/delete/product/:product_id",ProductController.RemoveProduct)
	server.DELETE("/delte/user/:user_id", UserController.RemoveUser)





	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}