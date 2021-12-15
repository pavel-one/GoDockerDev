package routes

import (
	"app/controller/IndexController"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", IndexController.Index)
}
