package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.comElcioAngelo/TrabalhoTCC.git/model"
	"github.comElcioAngelo/TrabalhoTCC.git/usecase"
	"strconv"
)

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) userController {
	return userController{
		userUseCase: usecase,
	}
}

func (u *userController) GetUser(ctx *gin.Context) {
	user_id := ctx.Param("user_id")


	user, err := u.userUseCase.GetUser(user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *userController) CreateUser(ctx *gin.Context) {
	var user model.User



	// * Recebe o Body do JSON (nome,email..etc)
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	

	Usererror := u.userUseCase.CreateUser(user)
	if Usererror != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error": err.Error(),
		})
		return
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"message": "User created successfully!",
			"user_name": user.Name,
		})
		return 
	}
}

func (u *userController) RemoveUser(ctx *gin.Context) {
	id := ctx.Param("user_id")

	num, stringError := strconv.Atoi(id)
	if stringError != nil {
		ctx.JSON(500,gin.H{
			"message": "an error ocurred",
			"error": stringError,
		})
	}

	err := u.userUseCase.RemoveUser(num)
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot remove user",
			"error": err,
		})
	}
	ctx.JSON(200,gin.H{
		"message": "Successfully deleted user",
	})
}

