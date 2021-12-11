package core

import (
	"app/db"
	"github.com/joho/godotenv"
	"log"
)

type App struct {
	DB db.DB
}

func Load() App {
	//Init env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	app := App{}
	app.DB = db.Load()

	return app
}
