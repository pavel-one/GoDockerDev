package UserService

import (
	"app/core"
	"app/models"
	"github.com/brianvoe/gofakeit/v6"
)

func CreateTestUser() {

	u := models.User{
		Username: gofakeit.Username(),
		Name:     gofakeit.FirstName(),
	}

	core.GetAppInstance().DB.Orm.Create(&u)
}
