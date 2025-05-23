package main

import (
	"go-api/config"
	"go-api/router"

	"go-api/inject"
)

func main() {
	config.Load()
	init := inject.Init()
	app := router.Init(init)

	app.Run(":8000")
}
