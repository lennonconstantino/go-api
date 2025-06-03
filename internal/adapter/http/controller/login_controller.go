package controller

import (
	entity "go-api/internal/core/domain"
	"go-api/internal/core/usecase"

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

// @Summary Login User
// @Description Login user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} response.JSONSuccessResult{data=entity.Auth,code=int,message=string}
// @Failure 400 {object} response.JSONBadRequestResult{code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{code=int,message=string}
// @Router /api/login [post]
func (lc LoginControllerImpl) Login(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var u entity.User
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

	ctx.JSON(http.StatusOK, entity.Auth{ID: userID, Token: token})
}
