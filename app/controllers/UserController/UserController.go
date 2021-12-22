package UserController

import (
	"app/base"
	"app/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *base.DB
}

func New(db *base.DB) *UserController {
	return &UserController{db}
}

func (c UserController) Index(ctx *gin.Context) {
	var users []models.User

	c.DB.Find(&users)

	ctx.JSON(200, &users)
}

func (c UserController) Create(ctx *gin.Context) {
	var u = models.User{
		Username: gofakeit.Username(),
		Name:     gofakeit.Person().FirstName,
	}
	c.DB.Create(&u)

	ctx.JSON(201, &u)
}
