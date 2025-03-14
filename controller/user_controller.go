package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.comElcioAngelo/TrabalhoTCC.git/usecase"
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

	user, err := u.userUseCase.GetUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}

	ctx.JSON(http.StatusOK, user)
}
