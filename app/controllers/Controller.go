package controllers

import (
	"app/base"
	"app/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	DB *base.DB
}

func (c Controller) Error(errors map[string]interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"errors":  errors,
	})
}

func (c Controller) Success(model interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success":  true,
		"resource": model,
	})
}

func (c Controller) AuthMiddleware(ctx *gin.Context) {
	tokenString, err := helpers.GetToken(ctx)
	if err != nil {
		ctx.AbortWithStatus(403)
		return
	}

	user, err := helpers.GetUserWithToken(tokenString, c.DB)
	if err != nil {
		ctx.AbortWithStatus(403)
		return
	}

	user.Token.UpdateUse(c.DB)

	ctx.Next()
}
