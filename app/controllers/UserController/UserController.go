package UserController

import (
	"app/base"
	"app/models"
	"encoding/json"
	"fmt"
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

	j, err := json.Marshal(users)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Нет объектов для выборки",
		})
	}

	fmt.Println(j)

	ctx.JSON(200, gin.H{})
}

func (c UserController) Create(ctx *gin.Context) {
	var u = models.User{
		Username: gofakeit.Username(),
		Name:     gofakeit.Person().FirstName,
	}
	c.DB.Create(&u)
}
