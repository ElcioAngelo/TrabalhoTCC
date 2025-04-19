package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"trabalhoTcc.com/mod/middleware"
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
	// Get user ID from URL parameter
	requestedUserID := ctx.Param("user_id")
	
	// Get user ID from cookie
	userIDValue, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	cookieUserID := fmt.Sprintf("%v", userIDValue)
	// Check if the cookie user ID matches the one in the path
	if cookieUserID != requestedUserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not permitted to view or modify other users",
		})
		return
	}

	// Fetch the user from the repository
	user, err := u.repository.GetUserByID(requestedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the user data
	ctx.JSON(http.StatusOK, user)
}


func (u *userController) UserVerification(ctx *gin.Context) {

	type request struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	

	var userInfo request 

	err := ctx.ShouldBind(&userInfo);
	if err != nil {
		ctx.JSON(400,gin.H{
			"message": "error while binding user information.",
			"error": err.Error(),
		});
		return
	}

	user, err := u.repository.UserVerification(userInfo.Email,userInfo.Password)
	if err != nil {
		ctx.JSON(404,gin.H{
			"message": "user not found",
			"error": err.Error(),
		});
		return
	}

	tokenString, err := middleware.GenerateToken(user.ID,user.Email,user.UserRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create token"})
		return
	}
	
	cookie := &http.Cookie{
		Name:     "jwtToken",
		Value:    tokenString,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   60 * 60 * 24 * 2, // 2 days
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Secure:   false, // Set to true in production (with HTTPS)
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	fmt.Printf("user information: %d, %s, %s",user.ID, user.Email, user.UserRole)
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
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message": "User created successfully!",
		"user_name": user.Name,
	})
}



func (u *userController) RemoveUser(ctx *gin.Context) {
	type status struct {
		Value string  `json:"value"`
	}

	// TODO: user remover.

}

func (u *userController) UserVerifyID(user_id string, ctx *gin.Context) (bool){

	user, err := u.repository.GetUserByID(user_id)
	if err != nil {
		ctx.JSON(404,gin.H{
			"message": "user not found",
			"error": err.Error(),
		});
		return false
	}

	id, Iderr := strconv.Atoi(user_id);
	if Iderr != nil {
		ctx.JSON(500,gin.H{
			"message": "error while transforming id to int",
			"error": Iderr.Error(),
		})
		return false
	}

	if (user.ID != id) { 
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"message": "user is not authorized to update",
		})
		return false
	}
	return true
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


// func(u *userController) setUserOrder(ctx *gin.Context) {
// 	user_id := 
// 	id, err := strconv.Atoi(user_id); if err != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{
// 			"message": "error while parsing int",
// 		})
// 	}
	
// } 

func(u *userController) AuthMeHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	email := c.GetString("email")
	userRole := c.GetString("user_role")

	fmt.Printf("user id: %s, user email: %s, user role: %s", userID, email, userRole);

	c.JSON(http.StatusOK, gin.H{
		"user_id":   userID,
		"email":     email,
		"user_role": userRole,
	})
}

