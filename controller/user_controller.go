package controller

import (
	"encoding/json"
	"errors"
	"go-api/controller/authentication"
	"go-api/model"
	"go-api/security"
	"go-api/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserControler interface {
	GetUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}

type UserControllerImpl struct {
	userUsecase usecase.UserUsecase
}

// NewUserController initialize
func NewUserController(usecase usecase.UserUsecase) *UserControllerImpl {
	return &UserControllerImpl{
		userUsecase: usecase,
	}
}

func (uu UserControllerImpl) GetUsers(ctx *gin.Context) {
	username := strings.ToLower(ctx.Query("username"))

	users, err := uu.userUsecase.GetUsers(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uu UserControllerImpl) GetUserById(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := uu.userUsecase.GetUserById(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uu UserControllerImpl) CreateUser(ctx *gin.Context) {

	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	if err = user.Prepare("form"); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err = uu.userUsecase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (uu UserControllerImpl) UpdateUser(ctx *gin.Context) {
	userIDToken, err := authentication.ExtractIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, userIDToken)
		return
	}

	userID, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if userID != userIDToken {
		ctx.JSON(http.StatusForbidden, errors.New("It is not possible to update a user that is not yours"))
		return
	}

	bodyRequest, err := ctx.GetRawData()
	//fmt.Printf("The user request value %v", bodyRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("edit"); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := uu.userUsecase.UpdateUser(userID, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (uu UserControllerImpl) DeleteUser(ctx *gin.Context) {
	userIDToken, err := authentication.ExtractIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, userIDToken)
		return
	}

	userID, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if userID != userIDToken {
		ctx.JSON(http.StatusForbidden, errors.New("It is not possible to delete a user that is not yours."))
		return
	}

	if err := uu.userUsecase.DeleteUser(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (uu UserControllerImpl) UpdatePassword(ctx *gin.Context) {
	userIDToken, err := authentication.ExtractIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, userIDToken)
		return
	}

	userID, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if userIDToken != userID {
		ctx.JSON(http.StatusForbidden, errors.New("It is not possible to change a user that is not yours"))
		return
	}

	bodyRequest, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	var password model.Password
	if err = json.Unmarshal(bodyRequest, &password); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	passwordInDatabase, err := uu.userUsecase.FetchPassword(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(passwordInDatabase, password.Current); err != nil {
		ctx.JSON(http.StatusUnauthorized, errors.New("The current password does not match the one saved in the database"))
		return
	}

	passwordWithHash, err := security.Hash(password.New)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err = uu.userUsecase.UpdatePassword(userID, string(passwordWithHash)); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
