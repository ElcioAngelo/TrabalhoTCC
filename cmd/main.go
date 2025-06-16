package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"trabalhoTcc.com/mod/controller"
	"trabalhoTcc.com/mod/database"
	"trabalhoTcc.com/mod/middleware"
	"trabalhoTcc.com/mod/repository"
)

func main() {
	server := gin.Default()
	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	// ** Configuração CORS para ambos os servidores conseguirem se comunicar
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// ** Definição de rotas protegidas
	protected := server.Group("/auth") // ** Rotas protegidas contem "/auth" na url.
	protected.Use(middleware.TokenAuthMiddleware())

	// ** Repositório e controladores (Pedidos e usuários)

	UserRepository := repository.NewUserRepository(dbConnection)
	UserController := controller.NewUserController(UserRepository)

	OrderRepository := repository.NewOrderRepository(dbConnection)
	OrderController := controller.NewOrderController(OrderRepository)

	// * Repositório e controladores (Produtos)

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductController := controller.NewProductController(ProductRepository)

	// * Requisições GET

	server.GET("/current_user", UserController.AuthMeHandler)
	protected.GET("/fetch/user/:user_id", UserController.GetUser)
	server.GET("/fetch/products", ProductController.GetProducts)
	protected.GET("/fetch/products/all", ProductController.GetProductsAdmin)
	server.GET("/fetch/product/:product_id", ProductController.GetProduct)
	server.GET("/fetch/product/by/:category", ProductController.FindProductByCategory)
	protected.GET("/fetch/all/users", UserController.GetUsers)
	protected.GET("/fetch/user/order", OrderController.ReturnOrder)
	protected.GET("/fetch/all/users/orders", OrderController.ReturnAllOrders)

	// * Requisições POST

	server.POST("/create/product", ProductController.CreateProduct)
	server.POST("/create/user", UserController.CreateUser)
	server.POST("/user/login", UserController.UserVerification)
	protected.POST("/make_user/order", OrderController.SetUserOrder)

	protected.PATCH("/edit/user", UserController.UpdateUser)

	// * Verificação do funcionamento do servidor
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}
