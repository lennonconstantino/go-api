//go:build wireinject
// +build wireinject

package inject

import (
	"go-api/internal/adapter/http/controller"
	"go-api/internal/adapter/repository"
	"go-api/internal/adapter/repository/postgres"
	"go-api/internal/core/usecase"

	"github.com/google/wire"
)

var pqconn = wire.NewSet(postgres.Connect)

var userRepositorySet = wire.NewSet(postgres.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*postgres.UserRepositoryImpl)),
)

var productRepositorySet = wire.NewSet(postgres.NewProductRepository,
	wire.Bind(new(repository.ProductRepository), new(*postgres.ProductRepositoryImpl)),
)

/*
var redisconn = wire.NewSet(redis.RedisConnect)

var cacheRepositorySet = wire.NewSet(redis.NewCacheRepository,
	wire.Bind(new(redis.CacheRepository), new(*redis.CacheRepositoryImpl)),
)

var userCacheRepositorySet = wire.NewSet(redis.NewUserCacheRepository,
	wire.Bind(new(redis.UserCacheRepository), new(*redis.UserCacheRepositoryImpl)),
	//wire.Bind(new(repository.UserRepository), new(*postgres.UserRepositoryImpl)),
)

var productCacheRepositorySet = wire.NewSet(redis.NewProductCacheRepository,
	wire.Bind(new(redis.ProductCacheRepository), new(*redis.ProductCacheRepositoryImpl)),
	//wire.Bind(new(repository.ProductRepository), new(*postgres.ProductRepositoryImpl)),
)
*/

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
	wire.Build(NewInitialization, pqconn, userRepositorySet, productRepositorySet, userUsecaseSet, productUsecaseSet, loginControllerSet, userControllerSet, productControllerSet)
	return nil
}
