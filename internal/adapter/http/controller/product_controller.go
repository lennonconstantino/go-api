package controller

import (
	"encoding/json"
	"fmt"
	"go-api/internal/adapter/repository/redis"
	entity "go-api/internal/core/domain"
	"go-api/internal/core/usecase"
	"go-api/utils"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetProducts(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

// ProductController struct
type ProductControllerImpl struct {
	productUsecase usecase.ProductUsecase
}

// NewProductController initialize
func NewProductController(usecase usecase.ProductUsecase) *ProductControllerImpl {
	return &ProductControllerImpl{
		productUsecase: usecase,
	}
}

// GetProdutcts godoc
// @Summary GetProdutcts
// @Description GetProdutcts
// @Accept  json
// @Produce  json
// @Tags Products
// @Success 200 {object} response.JSONSuccessResult{data=[]entity.Product,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/products [get]
func (p ProductControllerImpl) GetProducts(ctx *gin.Context) {
	// Accessing a header using c.GetHeader method
	contentType := ctx.GetHeader("cache")
	var products []entity.Product
	var err error

	if contentType == "true" {
		cacheRepository := redis.NewCacheRepository(redis.RedisConnect())
		reply, err := cacheRepository.Get("products")
		if err != nil {
			products, err = p.productUsecase.GetProducts()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			productBytes, _ := json.Marshal(products)
			cacheRepository.Set("products", productBytes, nil)
		}

		json.Unmarshal(reply, &products)
		ctx.JSON(http.StatusOK, products)
		return
	}

	fmt.Println("No Cache")
	products, err = p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary CreateProduct
// @Description CreateProduct
// @Accept  json
// @Produce  json
// @Tags Products
// @Param product body dto.ProductCreateRequestBody true "Product Data"
// @Success 200 {object} response.JSONSuccessResult{data=entity.Product,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/product [post]
func (p ProductControllerImpl) CreateProduct(ctx *gin.Context) {
	userId, err := utils.ExtractIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, userId)
		return
	}

	var product entity.Product
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	// product.Prepare and Validations - Usecase?

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

// GetProductById godoc
// @Summary GetProductById
// @Description GetProductById
// @Tags Products
// @Accept  json
// @Produce  json
// @Param productId   path   int true   "ProductRequestParam"
// @Success 200 {object} response.JSONSuccessResult{data=entity.Product,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/product/{productId} [get]
func (p ProductControllerImpl) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := entity.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.Response{
			Message: "Id do produto precisa ser numerico",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	contentType := ctx.GetHeader("cache")
	var product *entity.Product

	if contentType == "true" {
		cacheRepository := redis.NewCacheRepository(redis.RedisConnect())
		reply, err := cacheRepository.Get(fmt.Sprintf("products:%d", productId))
		if err != nil {
			product, err = p.productUsecase.GetProductById(productId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if product == nil {
				response := entity.Response{
					Message: "Produto nao foi encontrado na base de dados",
				}
				ctx.JSON(http.StatusNotFound, response)
				return
			}
			productBytes, _ := json.Marshal(product)
			cacheRepository.Set(fmt.Sprintf("products:%d", productId), productBytes, nil)
		}

		json.Unmarshal(reply, &product)
		ctx.JSON(http.StatusOK, product)
		return
	}

	product, err = p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := entity.Response{
			Message: "Produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary DeleteProduct
// @Description DeleteProduct
// @ID productId
// @Accept  json
// @Produce  json
// @Tags Products
// @Param productId   path   int true   "ProductRequestParam"
// @Success 204 {object} response.JSONSuccessResult{data=nil,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/product/{productId} [delete]
func (p ProductControllerImpl) DeleteProduct(ctx *gin.Context) {
	userId, err := utils.ExtractIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, userId)
		return
	}

	id := ctx.Param("productId")
	if id == "" {
		response := entity.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.Response{
			Message: "Id do produto precisa ser numerico",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := entity.Response{
			Message: "Produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if err = p.productUsecase.DeleteProduct(productId); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UpdateProduct godoc
// @Summary UpdateProduct
// @Description UpdateProduct
// @ID productId
// @Accept  json
// @Produce  json
// @Tags Products
// @Param productId   path   int true   "ProductRequestParam"
// @Param product body dto.ProductCreateRequestBody true "Product Data"
// @Success 204 {object} response.JSONSuccessResult{data=nil,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/product/{productId} [put]
func (p ProductControllerImpl) UpdateProduct(ctx *gin.Context) {
	userId, err := utils.ExtractIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, userId)
		return
	}

	id := ctx.Param("productId")
	if id == "" {
		response := entity.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.Response{
			Message: "Id do produto precisa ser numerico",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	bodyRequest, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	var product entity.Product
	if err = json.Unmarshal(bodyRequest, &product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err = p.productUsecase.UpdateProduct(productId, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
