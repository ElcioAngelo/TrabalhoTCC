package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"trabalhoTcc.com/mod/controller"
	"trabalhoTcc.com/mod/db"
	"trabalhoTcc.com/mod/middleware"
	"trabalhoTcc.com/mod/repository"
)



func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	
	server.Use(cors.New(cors.Config{
		// Allow requests from your front-end
		AllowOrigins: []string{"http://localhost:3000"}, // Replace with your front-end origin
		// Allow headers and methods as needed
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	protected := server.Group("/auth")
	protected.Use(middleware.TokenAuthMiddleware())

	// !! Rotas chamadas "servers" são rotas desprotegidas, não exigem JSONWEBTOKEN,
	// !! Rotas chamadas "protected" são rotas protegidas, exigem JSONWEBTOKEN.
	
	// * Usuários ################################################################################################
	
	UserRepository := repository.NewUserRepository(dbConnection)
	UserController := controller.NewUserController(UserRepository)


	// * Produtos ################################################################################################
	
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductController := controller.NewProductController(ProductRepository)

	// * Requisições GET ################################################################################################
	protected.GET("/current_user",UserController.AuthMeHandler)
	protected.GET("/fetch/user/:user_id",UserController.GetUser)
	server.GET("/fetch/products", ProductController.GetProducts)
	protected.GET("/fetch/products/all",ProductController.GetProductsAdmin)
	server.GET("/fetch/product/:product_id",ProductController.GetProduct)
	
	// * Requisições POST ################################################################################################
		
	server.POST("/create/product",ProductController.CreateProduct)
	server.POST("/create/user/",UserController.CreateUser)
	server.POST("/user/login", UserController.UserVerification)
	
	// * Requisições PUT ################################################################################################
	
	protected.PUT("/edit/user/:user_id/name",UserController.EditUserName)
	protected.PUT("/edit/user/:user_id/email",UserController.EditUserEmail)
	protected.PUT("/edit/user/:user_id/password",UserController.EditUserPassword)
	protected.PUT("/edit/user/:user_id/paymentaddress",UserController.EditUserPaymentAdress)
	protected.PUT("/edit/user/:user_id/shipmentaddress",UserController.EditUserShipmentAdress)
	protected.PUT("/edit/user/:user_id/cellphone",UserController.EditUserCellphone)

	protected.PUT("/edit/Product/:product_id/name",ProductController.EditProductName)
	protected.PUT("/edit/Product/:product_id/price",ProductController.EditProductPrice)
	protected.PUT("/edit/Product/:product_id/description",ProductController.EditProductDescription)
	
	// * Requisições DELETE ################################################################################################

	protected.DELETE("/delte/user/:user_id", UserController.RemoveUser)

	// * Verificação do funcionamento do servidor
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}