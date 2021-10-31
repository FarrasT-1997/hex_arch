package controller

import (
	"hex/controller/request"
	"hex/controller/response"
	"hex/database/port"
	"hex/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type AuthorController struct {
	authorModel port.AuthorPort
}

func NewController(authorModel port.AuthorPort) *AuthorController {
	return &AuthorController{
		authorModel: authorModel,
	}
}

func (controller *AuthorController) GetAllAuthor(c echo.Context) error {
	author, err := controller.authorModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, author)
}

func (controller *AuthorController) GetAuthorById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	author, err := controller.authorModel.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, author)
}

func (controller *AuthorController) PostAuthorController(c echo.Context) error {
	// bind request value
	var authorRequest request.PostAuthorRequest

	if err := c.Bind(&authorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	author := models.Author{
		Name:     authorRequest.Name,
		Email:    authorRequest.Email,
		Sex:      authorRequest.Sex,
		Password: authorRequest.Password,
	}

	_, err := controller.authorModel.Insert(author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.InternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, response.SuccessResponse())
}

func (controller *AuthorController) EditAuthorController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	// bind request value
	var authorRequest request.EditAuthorRequest
	if err := c.Bind(&authorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	author := models.Author{
		Name:     authorRequest.Name,
		Email:    authorRequest.Email,
		Sex:      authorRequest.Sex,
		Password: authorRequest.Password,
	}

	if _, err := controller.authorModel.Edit(author, id); err != nil {
		return c.JSON(http.StatusNotFound, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, response.SuccessResponse())
}

func (controller *AuthorController) DeleteAuthorController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	if _, err := controller.authorModel.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, response.SuccessResponse())
}

func (controller *AuthorController) LoginAuthorController(c echo.Context) error {
	var authorRequest request.LoginAuthorRequest

	if err := c.Bind(&authorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	author, err := controller.authorModel.Login(authorRequest.Email, authorRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": author.Token,
	})
}
