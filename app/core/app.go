//TODO: Сделать нормальную инциализацию роутов через роутер
//TODO: Сделать нормальную инциализацию контроллеров через new

package core

import (
	"app/base"
	"github.com/joho/godotenv"
	"log"
)

type App struct {
	DB     *base.DB
	Router base.Router
}

func NewApp() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	return &App{
		DB:     base.LoadDB(),
		Router: base.LoadRouter(),
	}
}

func (a App) Run() {
	err := a.Router.Run(":8080")

	if err != nil {
		log.Fatal("Ошибка запуска сервера " + err.Error())
	}
}
