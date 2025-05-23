package inject

import (
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"
)

type Initialization struct {
	UserRepository    repository.UserRepository
	ProductRepository repository.ProductRepository
	UserUsecase       usecase.UserUsecase
	ProductUsecase    usecase.ProductUsecase
	LoginController   controller.LoginController
	UserController    controller.UserController
	ProductController controller.ProductController
}

func NewInitialization(userRepository repository.UserRepository,
	productRepository repository.ProductRepository,
	userUsecase usecase.UserUsecase,
	productUsecase usecase.ProductUsecase,
	loginController controller.LoginController,
	userController controller.UserController,
	productController controller.ProductController) *Initialization {
	return &Initialization{
		UserRepository:    userRepository,
		ProductRepository: productRepository,
		UserUsecase:       userUsecase,
		ProductUsecase:    productUsecase,
		LoginController:   loginController,
		UserController:    userController,
		ProductController: productController,
	}
}
