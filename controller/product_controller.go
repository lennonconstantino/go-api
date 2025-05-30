package controller

import (
	"encoding/json"
	"go-api/model"
	"go-api/usecase"
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

// GetProducts get products
func (p ProductControllerImpl) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

// Create Product in database
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

// GetProductById
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

// DeleteProduct
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

// UpdateProduct
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
