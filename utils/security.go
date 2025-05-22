package utils

import (
	"errors"
	"fmt"
	"go-api/config"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Hash receive a string and hash it
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerifyPassword compare a password and a hash and return whether they are equal
func VerifyPassword(senhaString, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaString), []byte(senhaComHash))
}

// returnVerificationKey
func returnVerificationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

// CreateToken
func CreateToken(id uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"id":         id,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ExtractToken
func ExtractToken(ctx *gin.Context) string {
	ctx.Header("Content-Type", "application/json")
	tokenString := ctx.GetHeader("Authorization")

	if tokenString != "" {
		return tokenString[len("Bearer "):]
	}

	return ""
}

// ExtractIDFromToken
func ExtractIDFromToken(ctx *gin.Context) (uint64, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["id"]), 10, 64)
		if erro != nil {
			return 0, nil
		}

		return userID, nil
	}

	return 0, errors.New("Invalid Token")
}

// ValidateToken
func ValidateToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(ctx)
	if tokenString == "" {
		return nil, errors.New("Missing authorization header")
	}

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
