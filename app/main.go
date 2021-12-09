package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pioz/faker"
	"log"
	"os"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}

func main() {
	for {
		fmt.Println(faker.FirstName() + " " + faker.LastName())
		fmt.Println(os.Getenv("ENV_TEST") + "\n")

		time.Sleep(time.Second * 5)
	}
}
