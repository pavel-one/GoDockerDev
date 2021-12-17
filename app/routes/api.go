package routes

import (
	"app/controllers/UserController"
)

func (r *Route) Api() {
	var userController = UserController.New(r.DB)

	var api = r.Router.Group("api")
	{
		api.GET("/", userController.Index)
		api.GET("/create", userController.Create)
	}
}
