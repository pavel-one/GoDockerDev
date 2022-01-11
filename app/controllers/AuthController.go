package controllers

import (
	"app/base"
	"app/models"
	"app/requests"
	"app/resources"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	*Controller
}

func NewAuth(db *base.DB) *AuthController {
	controller := Controller{DB: db}

	return &AuthController{&controller}
}

func (c AuthController) Auth(ctx *gin.Context) {
	var request requests.Auth
	var user models.User

	if !requests.Validate(&request, ctx) {
		return
	}

	c.DB.Model(&user).Where("email = ?", request.Email).First(&user)

	if user.ID == 0 {
		c.Error(map[string]interface{}{
			"email": "Не правильное имя пользователя или пароль",
		}, ctx)
		return
	}

	if !user.CheckPasswordHash(request.Password) {
		c.Error(map[string]interface{}{
			"email": "Не правильное имя пользователя или пароль",
		}, ctx)
		return
	}
	user.CreateToken(c.DB).SetActualToken(c.DB)
	resource := resources.GetUserResource(&user)

	c.Success(resource, ctx)
}

func (c AuthController) Register(ctx *gin.Context) {
	var request requests.Register

	if !requests.Validate(&request, ctx) {
		return
	}

	var u = models.User{
		Username: request.Email,
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}

	_, err := u.Save(c.DB)

	if err != nil {
		c.Error(map[string]interface{}{
			"email": err.Error(),
		}, ctx)
		return
	}

	resource := resources.GetUserResource(&u)

	c.Success(resource, ctx)
}
