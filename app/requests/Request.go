package requests

import (
	"app/exceptions/ValidationExeption"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate(r interface{}, ctx *gin.Context) bool {
	err := ctx.ShouldBindJSON(&r)

	if err == nil {
		return true
	}

	e := ValidationExeption.New(err)
	ctx.JSON(http.StatusBadRequest, gin.H{"errors": e.FormatToFront()})
	return false
}
