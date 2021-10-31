package main

import (
	"fmt"
	"hex/config/connectDB"
	"hex/config/setting"
	"hex/controller"
	"hex/database/core"
	"hex/route"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	config := setting.GetConfig()

	db := connectDB.Connection(config)

	newAuthorModel := core.NewAuthorModel(db)

	newAuthorController := controller.NewController(newAuthorModel)

	e := echo.New()

	allController := route.Route{
		AuthorController: newAuthorController,
	}

	route.AllPath(e, allController)

	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
