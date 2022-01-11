package main

import (
	"app/core"
	_ "github.com/brianvoe/gofakeit/v6"
)

func main() {
	var app = core.NewApp()
	app.LoadRoutes(app.Route.Auth, app.Route.User)

	app.Run()
}
