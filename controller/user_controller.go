package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.repository.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while fetching users",
			"error":   err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *userController) GetUser(ctx *gin.Context) {
	requestedUserIDString := ctx.Param("user_id")

	requestedUserID, err := strconv.Atoi(requestedUserIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}

	userIDValue, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	if userIDValue != requestedUserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not permitted to view or modify other users",
		})
		return
	}

	user, err := u.repository.GetUserByID(requestedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(user)
	ctx.JSON(http.StatusOK, user)
}

func (u *userController) UserVerification(ctx *gin.Context) {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userInfo request

	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error while binding user information.",
			"error":   err.Error(),
		})
		return
	}

	user, err := u.repository.UserVerification(userInfo.Email, userInfo.Password)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "user not found",
			"error":   err.Error(),
		})
		return
	}

	tokenString, err := middleware.GenerateToken(user.ID, user.Email, user.UserRole)
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
		HttpOnly: false,
		Secure:   false, // Set to true in production (with HTTPS)
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
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
	if Usererr != nil {
		if pqError, ok := Usererr.(*pq.Error); ok {
			if pqError.Code == "23505" {
				ctx.JSON(409,gin.H{"Conflito": "Usuário já existente"})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
			return
	}else{
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "User created successfully!",
			"user_name": user.Name,
		})
	}
}



func (u *userController) RemoveUser(ctx *gin.Context) {
	type status struct {
		Value string `json:"value"`
	}

	// TODO: user remover.

}

func (u *userController) UserVerifyID(user_id int, ctx *gin.Context) bool {

	user, err := u.repository.GetUserByID(user_id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "user not found",
			"error":   err.Error(),
		})
		return false
	}

	if user.ID != user_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "user is not authorized to update",
		})
		return false
	}
	return true
}

func (u *userController) AuthMeHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	email := c.GetString("email")
	userRole := c.GetString("user_role")

	c.JSON(http.StatusOK, gin.H{
		"user_id":   userID,
		"email":     email,
		"user_role": userRole,
	})
}
func (u *userController) UpdateUser(c *gin.Context) {
	userID := c.GetInt("user_id")

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input", "details": err.Error()})
		return
	}

	if len(input) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data provided to update"})
		return
	}

	err := u.repository.UpdateUser(input, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User updated successfully"})
}
