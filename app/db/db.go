package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DB struct {
	Port     string
	Host     string
	User     string
	Password string
	Database string
	Dsn      string
	Orm      *gorm.DB
}

func (db *DB) Connect() {
	db.Dsn = "host=" + db.Host + " user=" + db.User + " password=" + db.Password + " dbname=" + db.Database + " port=" + db.Port + " sslmode=disable TimeZone=Europe/Moscow"
	connect, err := gorm.Open(postgres.Open(db.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД " + err.Error())
	}

	db.Orm = connect
}

func Load() DB {
	db := DB{}

	db.Port = os.Getenv("DB_PORT")
	db.Host = os.Getenv("DB_HOST")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Database = os.Getenv("DB_NAME")
	db.Connect()

	return db
}
