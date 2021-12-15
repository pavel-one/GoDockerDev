package main

import (
	"app/core"
)

func init() {
	core.Load()
}

func main() {
	core.GetAppInstance().Run()
}

func SuperTestFunc() string {
	return "TEST"
}
