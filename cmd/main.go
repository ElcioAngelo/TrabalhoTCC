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

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUsecase := usecase.NewUserRepository(UserRepository)
	UserController := controller.NewUserController(UserUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "pong",
		})
	})


	server.GET("/user",UserController.GetUser)
	








	server.Run(":8000")
	
}