package main

import (
	"encoding/json"
	"fmt"
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada de usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	// Camada de controllers
	productController := controller.NewProductController(ProductUsecase)

	server.POST("login", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")

		var u User
		json.NewDecoder(ctx.Request.Body).Decode(&u)
		fmt.Printf("The user request value %v", u)

		if u.Username == "Chek" && u.Password == "123456" {
			tokenString, err := createToken(u.Username)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				fmt.Errorf("No username found")
			}
			//ctx.JSON(http.StatusOK, nil)
			//fmt.Fprint(w, tokenString)
			//ctx.Writer.WriteString(tokenString)
			ctx.String(http.StatusOK, tokenString)
			return
		} else {
			ctx.String(http.StatusUnauthorized, "Invalid credentials")
		}
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)
	server.DELETE("/product/:productId", productController.DeleteProduct)
	server.PUT("/product/:productId", productController.UpdateProduct)

	server.Run(":8000")
}
