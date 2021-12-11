package db

import (
	"os"
)

type DB struct {
	Port     string
	Host     string
	User     string
	Password string
	Database string
	Dsn      string
}

func (db *DB) Connect() {
	db.Dsn = "host=" + db.Host + " user=" + db.User + " password=" + db.Password + " dbname=" + db.Database + " port=" + db.Port + " sslmode=disable TimeZone=Europe/Moscow"
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
