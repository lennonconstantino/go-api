package main

import (
	"go-api/config"
	"go-api/router"

	"go-api/inject"
)

func main() {

	config.Load()

	/*
		server := gin.Default()

		dbConnection, err := db.ConnectDB()
		if err != nil {
			panic(err)
		}

		userRepository := repository.NewUserRepository(dbConnection)
		userUsecase := usecase.NewUserUsecase(userRepository)
		userController := controller.NewUserController(userUsecase)

		loginController := controller.NewLoginController(userUsecase)

		productRepository := repository.NewProductRepository(dbConnection)
		productUsecase := usecase.NewProductUsecase(productRepository)
		productController := controller.NewProductController(productUsecase)

		server.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		public := server.Group("/api")
		public.POST("/login", loginController.Login)

		public.GET("/users", userController.GetUsers)
		public.GET("/user/:userId", userController.GetUserById)
		public.POST("/user", userController.CreateUser)
		public.GET("/products", productController.GetProducts)
		public.POST("/product", productController.CreateProduct)
		public.GET("/product/:productId", productController.GetProductById)

		protected := server.Group("/api/protected")
		protected.Use(middleware.JwtAuthMiddleware())
		protected.PUT("/user/:userId", userController.UpdateUser)
		protected.DELETE("/user/:userId", userController.DeleteUser)
		protected.POST("/user/:userId/update-password", userController.UpdatePassword)
		protected.DELETE("/product/:productId", productController.DeleteProduct)
		protected.PUT("/product/:productId", productController.UpdateProduct)

		server.Run(":8000")
	*/

	init := inject.Init()
	app := router.Init(init)

	app.Run(":8000")
}
