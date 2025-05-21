package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method! %v", token.Header["alg"])
	}

	//return config.SecretKey, nil
	return secretKey, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"username":   username, // change for id
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	//config.SecretKey
	return tokenString, nil
}

func ExtractToken(ctx *gin.Context) string {
	ctx.Header("Content-Type", "application/json")
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, "Missing authorization header")
		return ""
	}

	return tokenString[len("Bearer "):]
}

func ExtractUserName(ctx *gin.Context) (string, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := fmt.Sprintf("%s", permissions["username"])
		if username != "" {
			return username, nil
		}
	}

	return "", errors.New("Invalid Token")
}
