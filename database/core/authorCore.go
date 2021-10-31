package core

import (
	"hex/middlewares"
	"hex/models"

	"gorm.io/gorm"
)

type GormAuthorModel struct {
	db *gorm.DB
}

func NewAuthorModel(db *gorm.DB) *GormAuthorModel {
	return &GormAuthorModel{db: db}
}

func (m *GormAuthorModel) GetAll() ([]models.Author, error) {
	var authors []models.Author
	if err := m.db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (m *GormAuthorModel) Get(authorId int) (models.Author, error) {
	var author models.Author
	if err := m.db.Find(&author, authorId).Error; err != nil {
		return author, err
	}
	return author, nil
}

func (m *GormAuthorModel) Insert(author models.Author) (models.Author, error) {
	if err := m.db.Save(&author).Error; err != nil {
		return author, err
	}
	return author, nil
}

func (m *GormAuthorModel) Edit(newAuthor models.Author, authorId int) (models.Author, error) {
	var author models.Author
	if err := m.db.Find(&author, "id=?", authorId).Error; err != nil {
		return author, err
	}
	author.Name = newAuthor.Name
	author.Sex = newAuthor.Sex
	author.DoB = newAuthor.DoB
	author.Email = newAuthor.Email
	author.Password = newAuthor.Password

	if err := m.db.Save(&author).Error; err != nil {
		return author, err
	}
	return author, nil
}

func (m *GormAuthorModel) Delete(authorId int) (models.Author, error) {
	var author models.Author
	if err := m.db.Find(&author, "id=?", authorId).Error; err != nil {
		return author, err
	}
	if err := m.db.Delete(&author).Error; err != nil {
		return author, err
	}
	return author, nil
}

func (m *GormAuthorModel) Login(email, password string) (models.Author, error) {
	var author models.Author
	var err error

	if err = m.db.Where("email=? AND password=?", email, password).First(&author).Error; err != nil {
		return author, err
	}

	author.Token, err = middlewares.CreateToken(int(author.ID))

	if err != nil {
		return author, err
	}
	if err = m.db.Save(author).Error; err != nil {
		return author, err
	}
	return author, nil
}
