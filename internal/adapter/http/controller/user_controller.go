package controller

import (
	"encoding/json"
	"errors"
	entity "go-api/internal/core/domain"
	"go-api/internal/core/usecase"
	"go-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController interface {
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

// GetUsers godoc
// @Summary GetUsers
// @Description GetUsers
// @ID username
// @Accept  json
// @Produce  json
// @Param username   path   int true   "UserRequestParam"
// @Success 200 {object} response.JSONSuccessResult{data=[]entity.User,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/users [get]
func (uu UserControllerImpl) GetUsers(ctx *gin.Context) {
	username := strings.ToLower(ctx.Query("username"))

	users, err := uu.userUsecase.GetUsers(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetUserById godoc
// @Summary GetUserById
// @Description GetUserById
// @ID userId
// @Accept  json
// @Produce  json
// @Param userId   path   int true   "UserRequestParam"
// @Success 200 {object} response.JSONSuccessResult{data=entity.User,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/user/{userId} [get]
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

// CreateUser godoc
// @Summary CreateUser
// @Description CreateUser
// @Accept  json
// @Produce  json
// @Param user body dto.UserCreateRequestBody true "User Data"
// @Success 200 {object} response.JSONSuccessResult{data=entity.User,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/user [post]
func (uu UserControllerImpl) CreateUser(ctx *gin.Context) {

	var user entity.User
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

// UpdateUser godoc
// @Summary UpdateUser
// @Description UpdateUser
// @ID userId
// @Accept  json
// @Produce  json
// @Param userId   path   int true   "UserRequestParam"
// @Param user body dto.UserUpdateRequestBody true "User Data"
// @Success 204 {object} response.JSONSuccessResult{data=nil,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/user/{userId} [put]
func (uu UserControllerImpl) UpdateUser(ctx *gin.Context) {
	userIDToken, err := utils.ExtractIDFromToken(ctx)
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

	var user entity.User
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

// DeleteUser godoc
// @Summary DeleteUser
// @Description DeleteUser
// @ID userId
// @Accept  json
// @Produce  json
// @Param userId   path   int true   "UserRequestParam"
// @Success 204 {object} response.JSONSuccessResult{data=nil,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/user/{userId} [delete]
func (uu UserControllerImpl) DeleteUser(ctx *gin.Context) {
	userIDToken, err := utils.ExtractIDFromToken(ctx)
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

// UpdatePassword godoc
// @Summary UpdatePassword
// @Description UpdatePassword
// @Accept  json
// @Produce  json
// @Param userId   path   int true   "UserRequestParam"
// @Param login body dto.LoginRequestDto true "login Data"
// @Success 204 {object} response.JSONSuccessResult{data=nil,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/user/{userId}/update-password [post]
func (uu UserControllerImpl) UpdatePassword(ctx *gin.Context) {
	userIDToken, err := utils.ExtractIDFromToken(ctx)
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

	var password entity.Password
	if err = json.Unmarshal(bodyRequest, &password); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	passwordInDatabase, err := uu.userUsecase.FetchPassword(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err = utils.VerifyPassword(passwordInDatabase, password.Current); err != nil {
		ctx.JSON(http.StatusUnauthorized, errors.New("The current password does not match the one saved in the database"))
		return
	}

	passwordWithHash, err := utils.Hash(password.New)
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
