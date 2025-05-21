package controller

import (
	"encoding/json"
	"fmt"
	"go-api/controller/authentication"
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginController struct {
}

func NewLoginController() loginController {
	return loginController{}
}

func (lc *loginController) Login(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var u model.User
	json.NewDecoder(ctx.Request.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := authentication.CreateToken(u.Username)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			//fmt.Errorf("No username found")
		}
		ctx.String(http.StatusOK, tokenString)
		return
	} else {
		ctx.String(http.StatusUnauthorized, "Invalid credentials")
	}
}
