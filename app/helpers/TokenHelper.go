package helpers

import (
	"app/base"
	"app/models"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetToken(ctx *gin.Context) (string, error) {
	tokenString := ""

	auth := ctx.Request.Header["Authorization"]
	if auth == nil {
		return tokenString, errors.New("authorization header not found")
	}

	token := strings.Split(auth[0], " ")
	if token[0] != "Bearer" {
		return tokenString, errors.New("authorization not bearer")
	}

	tokenString = token[1]

	if tokenString == "" {
		return tokenString, errors.New("error token")
	}

	return tokenString, nil
}

func GetUserWithToken(token string, db *base.DB) (models.User, error) {
	var userToken models.UserToken
	var user models.User
	db.Model(&userToken).Where("token = ?", token).First(&userToken)

	if userToken.ID == 0 {
		return user, errors.New("token not valid")
	}

	user = userToken.GetUser(db)

	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	user.SetActualToken(db)

	return user, nil
}
