package main

import (
	"fmt"
	"go-api/config"
	"go-api/inject"
	"go-api/internal/adapter/http/router"
)

func main() {
	conf := config.GetConfig()
	init := inject.Init()
	app := router.Init(init)

	app.Run(fmt.Sprintf(":%d", conf.Server.Port))
}
