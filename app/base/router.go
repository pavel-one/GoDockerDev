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
	r.loadRoutes()

	return r
}

func (r Router) loadRoutes() {

}
