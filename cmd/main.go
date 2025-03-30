package main

import (
	"github.com/gin-gonic/gin"
	"trabalhoTcc.com/mod/controller"
	"trabalhoTcc.com/mod/db"
	"trabalhoTcc.com/mod/repository"
)



func main() {
	server := gin.Default()
	
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// * Usuários ################################################################################################
	
	UserRepository := repository.NewUserRepository(dbConnection)
	UserController := controller.NewUserController(UserRepository)


	// * Produtos ################################################################################################
	
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductController := controller.NewProductController(ProductRepository)

	// * Requisições GET ################################################################################################
	
	server.GET("/fetch/user/:user_id",UserController.GetUser)
	server.GET("/fetch/products", ProductController.GetProducts)
	server.GET("/fetch/products/all",ProductController.GetProductsAdmin)
	
	// * Requisições POST ################################################################################################
	
	// TODO: server.POST("/create/user", UserController.CreateUser)
	
	server.POST("/create/product",ProductController.CreateProduct)
	
	// * Requisições PUT ################################################################################################
	
	server.PUT("/edit/user/:user_id/name",UserController.EditUserName)
	server.PUT("/edit/user/:user_id/email",UserController.EditUserEmail)
	server.PUT("/edit/user/:user_id/password",UserController.EditUserPassword)
	server.PUT("/edit/user/:user_id/paymentaddress",UserController.EditUserPaymentAdress)
	server.PUT("/edit/user/:user_id/shipmentaddress",UserController.EditUserShipmentAdress)

	server.PUT("/edit/Product/:product_id/name",ProductController.EditProductName)
	server.PUT("/edit/Product/:product_id/price",ProductController.EditProductPrice)
	server.PUT("/edit/Product/:product_id/description",ProductController.EditProductDescription)
	
	// * Requisições DELETE ################################################################################################

	// todo: server.DELETE("/delete/product/:product_id",ProductController)
	server.DELETE("/delte/user/:user_id", UserController.RemoveUser)




	// * Verificação do funcionamento do servidor
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}