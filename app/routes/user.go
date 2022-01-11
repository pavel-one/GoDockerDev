package routes

import (
	"app/controllers"
)

func (r Route) User() {
	c := controllers.NewUser(r.DB)

	user := r.Router.Group("user").Use(c.AuthMiddleware)
	{
		user.GET("/", c.User)
	}
}
