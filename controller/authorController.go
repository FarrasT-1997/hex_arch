package controller

import (
	"hex/database/port"

	"github.com/labstack/echo"
)

type Controller struct {
	authorModel port.AuthorPort
}

func NewController(authorModel port.AuthorPort) *Controller {
	return &Controller{
		authorModel: authorModel,
	}
}

func (controller *Controller) GetAllAuthor(c echo.Context) error {
	author, err := controller.authorModel.GetAll()
}
