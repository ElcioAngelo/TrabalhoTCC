package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"trabalhoTcc.com/mod/model"
	"trabalhoTcc.com/mod/repository"
)

type userController struct {
	repository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) userController {
	return userController{
		repository: repository,
	}
}

func (u *userController) GetUser(ctx *gin.Context) {
	user_id := ctx.Param("user_id")


	user, err := u.repository.GetUser(user_id)
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

	

	Usererr := u.repository.CreateUser(user)
	if	Usererr != nil {
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
	type status struct {
		Value string  `json:"value"`
	}

	// TODO: user remover.

}

func (u *userController) EditUserName(ctx *gin.Context) {
	
	type UpdateUsername struct {
		Name string `json:"username"`
	}

	var request UpdateUsername
	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}

	err := ctx.ShouldBind(&request);
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot parse user name",
			"error": err.Error(),
		})
	}

	userUpdateError := u.repository.EditUserName(id,request.Name)
	if userUpdateError != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot change user name",
			"error": err.Error(),
		})
	}

	ctx.JSON(200,gin.H{
		"message": "user updated successfully!",
		"error": request.Name,
	})
}

func (u *userController) EditUserPassword(ctx *gin.Context) {
	
	type Update struct {
		Password string `json:"password"`
	}

	var request Update
	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}

	err := ctx.ShouldBind(&request);
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot parse user password",
			"error": err.Error(),
		})
	}

	userUpdateError := u.repository.EditUserPassword(id,request.Password)
	if userUpdateError != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot change user password",
			"error": err.Error(),
		})
	}

	ctx.JSON(200,gin.H{
		"message": "user password successfully!",
	})
}

func (u *userController) EditUserShipmentAdress(ctx *gin.Context) {
	
	type Update struct {
		ShippingAddress string `json:"shipping_address"`
	}

	var request Update
	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}

	err := ctx.ShouldBind(&request);
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot parse user password",
			"error": err.Error(),
		})
	}

	userUpdateError := u.repository.EditUserShipmentAdress(id,request.ShippingAddress)
	if userUpdateError != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot change user shipping address",
			"error": err.Error(),
		})
	}

	ctx.JSON(200,gin.H{
		"message": "user shipaddress successfully!",
	})
}

func (u *userController) EditUserPaymentAdress(ctx *gin.Context) {
	
	type Update struct {
		PaymentAddress string `json:"payment_address"`
	}

	var request Update
	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}

	err := ctx.ShouldBind(&request.PaymentAddress);
	if err != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot parse user payaddress",
			"error": err.Error(),
		})
	}

	userUpdateError := u.repository.EditUserPaymentAdress(id,request.PaymentAddress)
	if userUpdateError != nil {
		ctx.JSON(403,gin.H{
			"message": "cannot change user payaddress",
			"error": err.Error(),
		})
	}

	ctx.JSON(200,gin.H{
		"message": "user pay address successfully!",
	})
}

func(u *userController) EditUserCellphone(ctx *gin.Context) {
	type value struct {
		Cellphone string `json:"cellphone"`
	}

	var request value

	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}

	numberErr := ctx.ShouldBind(&request.Cellphone);
	if numberErr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while parsing user cellphone",
			"error": numberErr.Error(),
		})
	}
	
	err := u.repository.EditUserCellphone(id,request.Cellphone);
	if err != nil {
		ctx.JSON(500,gin.H{
			"message": "error while updating user cellphone",
			"error": numberErr.Error(),
		})
	}
}

func(u *userController) EditUserEmail(ctx *gin.Context) {
	type value struct {
		Email string `json:"email"`
	}

	var request value

	requestId := ctx.Param("user_id")
	id, Iderr := strconv.Atoi(requestId);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
	}

	numberErr := ctx.ShouldBind(&request.Email);
	if numberErr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while parsing user cellphone",
			"error": numberErr.Error(),
		})
	}
	
	err := u.repository.EditUserCellphone(id,request.Email);
	if err != nil {
		ctx.JSON(500,gin.H{
			"message": "error while updating user cellphone",
			"error": numberErr.Error(),
		})
	}
}