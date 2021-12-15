package core

import (
	"app/db"
	"app/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

type App struct {
	DB     db.DB
	Router *gin.Engine
}

var appInstance App

func Load() {
	//Init env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	app := App{}
	app.DB = db.Load()
	app.Router = gin.Default()
	routes.Routes(app.Router)

	appInstance = app
}

func GetAppInstance() *App {
	return &appInstance
}

func (a App) Run() {
	err := a.Router.Run(":8080")

	if err != nil {
		log.Fatal("Ошибка запуска сервера " + err.Error())
	}
}
