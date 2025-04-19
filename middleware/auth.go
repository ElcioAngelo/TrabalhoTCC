package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"trabalhoTcc.com/mod/model"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY")) // Signing key for JWT


// Middleware to validate JWT token
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		tokenString, err := c.Cookie("jwtToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{
				"message": "authorization token is missing",
			})
			c.Abort()
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
            return mySigningKey, nil
        })
		if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
            c.Abort()
            return
        }

		claims, ok := token.Claims.(*model.Claims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
            c.Abort()
            return
        }
		c.Set("email", claims.Email)
		c.Set("user_role",claims.UserRole)
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}




// Route to generate a JWT token
func GenerateToken(id int,email string, user_role string) (string, error){
	// Create the claims
	claims := model.Claims{
		UserID: id,
		Email: email,
		UserRole: user_role,
			RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "server",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 100)),
		},
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the signing key
	signedToken, err := token.SignedString(mySigningKey)

	// Return the signed token
	return signedToken, err
}


// Protected route that requires a valid JWT token
func ProtectedEndpoint(c *gin.Context) {
	// Access the username from the claims
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s! You have access to this protected route.", username)})
}

func VerifyUserToken(c *gin.Context, userIDParam int) (*model.Claims, bool) {
	// Get the token from the cookie
	tokenString, err := c.Cookie("jwtToken")
	if err != nil || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is missing"})
		return nil, false
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	// Handle errors in parsing or invalid token
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
		return nil, false
	}

	// Extract claims
	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
		return nil, false
	}

	// Check if the user ID in the token matches the user ID passed in the request (userIDParam)
	if claims.UserID != userIDParam {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID does not match the token"})
		return nil, false
	}

	return claims, true
}