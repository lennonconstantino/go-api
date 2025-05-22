package inject

import (
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"
)

type Initialization struct {
	userRepository    repository.UserRepository
	productRepository repository.ProductRepository
	userUsecase       usecase.UserUsecase
	productUsecase    usecase.ProductUsecase
	loginController   controller.LoginController
	userController    controller.UserController
	productController controller.ProductController
}

func NewInitialization(userRepository repository.UserRepository,
	productRepository repository.ProductRepository,
	userUsecase usecase.UserUsecase,
	productUsecase usecase.ProductUsecase,
	loginController controller.LoginController,
	userController controller.UserController,
	productController controller.ProductController) *Initialization {
	return &Initialization{
		userRepository:    userRepository,
		productRepository: productRepository,
		userUsecase:       userUsecase,
		productUsecase:    productUsecase,
		loginController:   loginController,
		userController:    userController,
		productController: productController,
	}
}
