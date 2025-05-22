package controller

import (
	"encoding/json"
	"go-api/controller/authentication"
	"go-api/model"
	"go-api/usecase"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// productController struct
type productController struct {
	ProductUsecase usecase.ProductUsecase
}

// NewProductController initialize
func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUsecase: usecase,
	}
}

// GetProducts get products
func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

// Create Product in database
func (p *productController) CreateProduct(ctx *gin.Context) {
	id, err := authentication.ExtractID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, id)
		return
	}

	var product model.Product
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	insertedProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

// GetProductById
func (p *productController) GetProductById(ctx *gin.Context) {
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

	product, err := p.ProductUsecase.GetProductById(productId)
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

func (p *productController) DeleteProduct(ctx *gin.Context) {
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

	// UseCase
	product, err := p.ProductUsecase.GetProductById(productId)
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

	if err = p.ProductUsecase.DeleteProduct(productId); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
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

	if err = p.ProductUsecase.UpdateProduct(productId, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
