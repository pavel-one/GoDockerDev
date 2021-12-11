package main

import (
	"app/core"
	"fmt"
	"time"
)

var App core.App

func init() {
	App = core.Load()
}

func main() {
	for {
		fmt.Println(App.DB.Dsn + "\n")

		time.Sleep(time.Second * 5)
	}
}
