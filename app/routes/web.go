package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (r Route) Web() {
	r.Router.GET("/", func(ctx *gin.Context) {
		fmt.Println("TEST")
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
