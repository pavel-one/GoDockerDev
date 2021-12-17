package controllers

import (
	"app/base"
	"github.com/gin-gonic/gin"
	"os"
)

type Controller struct {
	DB *base.DB
}

func (c Controller) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"version": os.Getenv("API_VERSION"),
		"mode":    os.Getenv("GIN_MODE"),
	})
}
