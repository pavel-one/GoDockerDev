package controllers

import (
	"app/base"
	"app/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	Index() string
	New()
}

type Controller struct {
	DB *base.DB
}

func (c Controller) Index(ctx *gin.Context) {
	var u = models.User{
		Username: gofakeit.Username(),
		Name:     gofakeit.Person().FirstName,
	}
	c.DB.Create(&u)

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
