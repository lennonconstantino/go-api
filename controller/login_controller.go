package controller

import (
	"go-api/model"
	"go-api/usecase"
	"go-api/utils"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type LoginControllerImpl struct {
	userUsecase usecase.UserUsecase
}

func NewLoginController(usecase usecase.UserUsecase) *LoginControllerImpl {
	return &LoginControllerImpl{
		userUsecase: usecase,
	}
}

func (lc LoginControllerImpl) Login(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var u model.User
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	userDB, err := lc.userUsecase.GetUserByEmail(u.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err = utils.VerifyPassword(userDB.Password, u.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	token, err := utils.CreateToken(userDB.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userDB.ID, 10)

	ctx.JSON(http.StatusOK, model.Auth{ID: userID, Token: token})
}
