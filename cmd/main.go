package main

import (
	"fmt"
	"go-api/config"
	"go-api/internal/adapter/http/router"

	"go-api/inject"
)

func main() {
	conf := config.GetConfig()
	init := inject.Init()
	app := router.Init(init)

	app.Run(":8000", fmt.Sprintf(":%d", conf.Server.Port))
}
