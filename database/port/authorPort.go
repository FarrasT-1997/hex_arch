package port

import "hex/models"

type AuthorPort interface {
	GetAll() ([]models.Author, error)
	Get(authorId int) (models.Author, error)
	Insert(models.Author) (models.Author, error)
	Edit(newAuthor models.Author, authorId int) (models.Author, error)
	Delete(authorId int) (models.Author, error)
	Login(email, password string) (models.Author, error)
}
