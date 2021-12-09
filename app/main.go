package main

import (
	"fmt"
	"github.com/joho/godotenv"
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
	var test = "string"
	fmt.Println(test)

	for {
		fmt.Println(os.Getenv("ENV_TEST") + "\n")

		time.Sleep(time.Second * 5)
	}
}
