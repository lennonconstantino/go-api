package router

import (
	"go-api/inject"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

func Init(init *inject.Initialization) *gin.Engine {
	//func Init() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// api := router.Group("/api")
	// {
	// 	user := api.Group("/user")
	// 	user.GET("", init.UserCtrl.GetAllUserData)
	// 	user.POST("", init.UserCtrl.AddUserData)
	// 	user.GET("/:userID", init.UserCtrl.GetUserById)
	// 	user.PUT("/:userID", init.UserCtrl.UpdateUserData)
	// 	user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	// }

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	public := router.Group("/api")
	public.POST("/login", init.LoginController.Login)
	public.GET("/users", init.UserController.GetUsers)
	public.GET("/user/:userId", init.UserController.GetUserById)
	public.POST("/user", init.UserController.CreateUser)
	public.GET("/products", init.ProductController.GetProducts)
	public.POST("/product", init.ProductController.CreateProduct)
	public.GET("/product/:productId", init.ProductController.GetProductById)

	protected := router.Group("/api/protected")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.PUT("/user/:userId", init.UserController.UpdateUser)
	protected.DELETE("/user/:userId", init.UserController.DeleteUser)
	protected.POST("/user/:userId/update-password", init.UserController.UpdatePassword)
	protected.DELETE("/product/:productId", init.ProductController.DeleteProduct)
	protected.PUT("/product/:productId", init.ProductController.UpdateProduct)

	return router
}
