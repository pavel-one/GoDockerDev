package main

import (
	"app/core"
	"app/routes"
	_ "github.com/brianvoe/gofakeit/v6"
)

func main() {
	var app = core.NewApp()
	routes.Web(app)

	app.Run()
}
