//go:build wireinject
// +build wireinject

package inject

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/google/wire"
)

var db2 = wire.NewSet(db.ConnectDB)

var userRepositorySet = wire.NewSet(repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var productRepositorySet = wire.NewSet(repository.NewProductRepository,
	wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)),
)

var userUsecaseSet = wire.NewSet(usecase.NewUserUsecase,
	wire.Bind(new(usecase.UserUsecase), new(*usecase.UserUsecaseImpl)),
)

var productUsecaseSet = wire.NewSet(usecase.NewProductUsecase,
	wire.Bind(new(usecase.ProductUsecase), new(*usecase.ProductUsecaseImpl)),
)

var loginControllerSet = wire.NewSet(controller.NewLoginController,
	wire.Bind(new(controller.LoginController), new(*controller.LoginControllerImpl)),
)

var userControllerSet = wire.NewSet(controller.NewUserController,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var productControllerSet = wire.NewSet(controller.NewProductController,
	wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db2, userRepositorySet, productRepositorySet, userUsecaseSet, productUsecaseSet, loginControllerSet, userControllerSet, productControllerSet)
	return nil
}
