package base

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func LoadRouter() Router {
	var r = Router{
		Engine: gin.Default(),
	}

	return r
}
