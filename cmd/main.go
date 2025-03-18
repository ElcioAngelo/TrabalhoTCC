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
	
	server.GET("/user/:user_id",UserController.GetUser)
	server.GET("/products", ProductController.GetProducts)
	
	// * Requisições POST ################################################################################################
	
	server.POST("/createUser", UserController.CreateUser)
	server.POST("/createProduct",ProductController.CreateProduct)
	
	// * Requisições PATCH ################################################################################################
	
	// !! Preciso de ajuda nas requisições de PATCH Para edição.

	
	// * Requisições DELETE ################################################################################################

	server.DELETE("/deleteProduct/:product_id",ProductController.RemoveProduct)
	server.DELETE("/user/:user_id", UserController.RemoveUser)





	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}