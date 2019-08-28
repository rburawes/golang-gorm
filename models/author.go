package models

import (
	"errors"
	"github.com/rburawes/golang-gorm/config"
)

// Author holds data about the author of the book.
type Author struct {
	AuthorID   int32  `json:"id" gorm:"primary_key"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	About      string `json:"about"`
}

// AllAuthors retrieve authors from the database.
func AllAuthors() ([]Author, error) {
	aus := make([]Author, 0)
	if config.Database.Find(&aus).Error != nil {
		return nil, errors.New("unable to find author records")
	}
	return aus, nil
}

// GetByID finds author record by id.
func GetByID(ids ...int32) ([]Author, error) {
	aus := make([]Author, 0)
	if config.Database.Where("author_id in (?)", ids).Find(&aus).Error != nil {
		return aus, errors.New("unable to find author using the given id")
	}

	return aus, nil
}
