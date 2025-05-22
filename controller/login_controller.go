package controller

import (
	"go-api/controller/authentication"
	"go-api/model"
	"go-api/security"
	"go-api/usecase"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	UserUsecase usecase.UserUsecase
}

func NewLoginController(usecase usecase.UserUsecase) loginController {
	return loginController{
		UserUsecase: usecase,
	}
}

func (lc *loginController) Login(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var u model.User
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	userDB, err := lc.UserUsecase.GetUserByEmail(u.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userDB.Password, u.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userDB.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userDB.ID, 10)

	ctx.JSON(http.StatusOK, model.Auth{ID: userID, Token: token})
}
