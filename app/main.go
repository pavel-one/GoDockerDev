package main

import (
	"app/core"
	_ "github.com/brianvoe/gofakeit/v6"
)

func main() {
	var app = core.NewApp()
	app.LoadRoutes(app.Route.Web, app.Route.Api)

	app.Run()
}
