package controller

import (
	"encoding/json"
	"fmt"
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

	//data, _ := ioutil.ReadAll(c.Request.Body)
	// _, err := ctx.GetRawData()
	// if err != nil {
	// 	ctx.JSON(http.StatusUnprocessableEntity, err)
	// 	return
	// }

	var u model.User
	json.NewDecoder(ctx.Request.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)

	userDB, err := lc.UserUsecase.GetUserByEmail(u.Email)
	if err != nil {
		fmt.Printf("1")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userDB.Password, u.Password); err != nil {
		fmt.Printf("2")
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userDB.ID)
	if err != nil {
		fmt.Printf("3")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userDB.ID, 10)

	ctx.JSON(http.StatusOK, model.Auth{ID: userID, Token: token})

	// if u.Username == "Chek" && u.Password == "123456" {
	// 	tokenString, err := authentication.CreateToken(u.ID)
	// 	if err != nil {
	// 		ctx.String(http.StatusInternalServerError, err.Error())
	// 		//fmt.Errorf("No username found")
	// 	}
	// 	ctx.String(http.StatusOK, tokenString)
	// 	return
	// } else {
	// 	ctx.String(http.StatusUnauthorized, "Invalid credentials")
	// }
}
