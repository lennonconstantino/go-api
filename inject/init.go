package inject

import (
	"go-api/internal/adapter/http/controller"
	"go-api/internal/adapter/repository"
	"go-api/internal/core/usecase"
)

type Initialization struct {
	UserRepository    repository.UserRepository
	ProductRepository repository.ProductRepository
	// CacheRepository        redis.CacheRepository
	// UserCacheRepository    redis.UserCacheRepository
	// ProductCacheRepository redis.ProductCacheRepository
	UserUsecase       usecase.UserUsecase
	ProductUsecase    usecase.ProductUsecase
	LoginController   controller.LoginController
	UserController    controller.UserController
	ProductController controller.ProductController
}

func NewInitialization(
	userRepository repository.UserRepository,
	productRepository repository.ProductRepository,
	// cacheRepository redis.CacheRepository,
	// userCacheRepository redis.UserCacheRepository,
	// productCacheRepository redis.ProductCacheRepository,
	userUsecase usecase.UserUsecase,
	productUsecase usecase.ProductUsecase,
	loginController controller.LoginController,
	userController controller.UserController,
	productController controller.ProductController) *Initialization {
	return &Initialization{
		//
		UserRepository:    userRepository,
		ProductRepository: productRepository,
		// CacheRepository:        cacheRepository,
		// UserCacheRepository:    userCacheRepository,
		// ProductCacheRepository: productCacheRepository,
		UserUsecase:       userUsecase,
		ProductUsecase:    productUsecase,
		LoginController:   loginController,
		UserController:    userController,
		ProductController: productController,
	}
}
