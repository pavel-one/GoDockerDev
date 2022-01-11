package routes

import (
	"app/controllers"
)

func (r Route) Auth() {
	c := controllers.NewAuth(r.DB)

	auth := r.Router.Group("auth")
	{
		auth.POST("register", c.Register)
		auth.POST("/", c.Auth)
	}
}
