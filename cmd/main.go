package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	UserRepository := repository.NewUserRepository(dbConnection)

	UserUsecase := usecase.NewUserUsecase(UserRepository)

	userController := controller.NewUserController(UserUsecase)

	loginController := controller.NewLoginController(UserUsecase)

	// camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada de usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	// Camada de controllers
	productController := controller.NewProductController(ProductUsecase)

	server.POST("login", loginController.Login)
	server.GET("protected", func(ctx *gin.Context) {
		// username, err := authentication.ExtractUserName(ctx)
		// if err != nil {
		// 	ctx.JSON(http.StatusUnauthorized, username)
		// 	return
		// }

		// ctx.JSON(http.StatusOK, "Welcome to the the protected area")
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/users", userController.GetUsers)
	server.GET("/user/:userId", userController.GetUserById)
	server.POST("/user", userController.CreateUser)
	server.PUT("/user/:userId", userController.UpdateUser)
	server.DELETE("/user/:userId", userController.DeleteUser)
	server.POST("/user/:userId/update-password", userController.UpdatePassword)

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)
	server.DELETE("/product/:productId", productController.DeleteProduct)
	server.PUT("/product/:productId", productController.UpdateProduct)

	server.Run(":8000")
}
