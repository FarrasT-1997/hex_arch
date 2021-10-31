package route

import (
	"hex/controller"

	"github.com/labstack/echo"
)

type Route struct {
	AuthorController *controller.AuthorController
}

func AllPath(e *echo.Echo, route Route) {
	e.POST("/author/register", route.AuthorController.PostAuthorController)
	e.POST("/author/login", route.AuthorController.LoginAuthorController)

	e.GET("/author", route.AuthorController.GetAllAuthor)
	e.GET("/author/:id", route.AuthorController.GetAuthorById)
	e.PUT("/author/:id", route.AuthorController.EditAuthorController)
	e.DELETE("/author/:id", route.AuthorController.DeleteAuthorController)
}
