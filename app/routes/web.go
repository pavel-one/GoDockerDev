package routes

import (
	"app/controllers"
	"app/core"
)

func Web(app *core.App) {
	app.Router.GET("/", controllers.Controller{DB: app.DB}.Index)
}
