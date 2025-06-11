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

var mySigningKey = []byte(os.Getenv("SECRET_KEY")) // ** Chave assinadora das JWT


// ** Middleware para a validação das JWT
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




// ** Rota de geração da JWT token
func GenerateToken(id int,email string, user_role string) (string, error){

	// ? Criando as "claims"
	claims := model.Claims{
		UserID: id,
		Email: email,
		UserRole: user_role,
			RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "server",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 100)),
		},
	}

	// ? Cria uma nova token com as "claims"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ? Assina a token com a chave de assinatura
	signedToken, err := token.SignedString(mySigningKey)

	// ? Retorna a token assinada
	return signedToken, err
}


// ** Rota protegida que requer uma JWT Token válida.
func ProtectedEndpoint(c *gin.Context) {
	// ? Acessando o nome do usuário nas "claims"
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s! You have access to this protected route.", username)})
}

func VerifyUserToken(c *gin.Context, userIDParam int) (*model.Claims, bool) {

	// ? Extração da token no cookie
	tokenString, err := c.Cookie("jwtToken")
	if err != nil || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is missing"})
		return nil, false
	}

	// ? "Parsando" a token.
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	// ? Validação de erros sobre tokens expiradas ou invalidas.
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
		return nil, false
	}

	// ? Extração da "claims"
	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
		return nil, false
	}

	// ? Verifica se o id do usuário na token é o mesmo do usuário atual
	if claims.UserID != userIDParam {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID does not match the token"})
		return nil, false
	}

	return claims, true
}