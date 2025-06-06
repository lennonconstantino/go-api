package controller

import (
	"encoding/json"
	"fmt"
	"go-api/internal/adapter/cache"
	model "go-api/internal/core/domain"
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
// @Success 200 {object} response.JSONSuccessResult{data=[]model.Product,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/products [get]
func (p ProductControllerImpl) GetProducts(ctx *gin.Context) {
	// Accessing a header using c.GetHeader method
	contentType := ctx.GetHeader("cache")
	var products []model.Product
	var err error

	if contentType == "true" {
		objects, err := cache.Cache("products", p.productUsecase.GetProducts)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		json.Unmarshal(objects, &products)
	} else {

		fmt.Println("No Cache")
		products, err = p.productUsecase.GetProducts()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
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
// @Success 200 {object} response.JSONSuccessResult{data=model.Product,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/product [post]
func (p ProductControllerImpl) CreateProduct(ctx *gin.Context) {
	userId, err := utils.ExtractIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, userId)
		return
	}

	var product model.Product
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
// @Success 200 {object} response.JSONSuccessResult{data=model.Product,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/product/{productId} [get]
func (p ProductControllerImpl) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
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
		response := model.Response{
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
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
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
		response := model.Response{
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
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
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

	var product model.Product
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
