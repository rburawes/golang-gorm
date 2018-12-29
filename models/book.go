package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rburawes/golang-demo/config"
	"net/http"
	"strconv"
)

// Book holds information about the book entry.
type Book struct {
	Isbn    string   `json:"isbn" gorm:"primary_key"`
	Title   string   `json:"title"`
	Price   float32  `json:"price"`
	Authors []Author `json:"authors" gorm:"many2many:book_authors;foreignkey:isbn;association_foreignkey:authorID;association_jointable_foreignkey:author_id;jointable_foreignkey:book_isbn"`
}

// AllBooks retrieve all the books from the database.
func AllBooks() ([]Book, error) {

	bks := make([]Book, 0)
	aus := make([]Author, 0)
	config.Database.Model(&bks).Related(&aus, "Authors")
	if config.Database.Preload("Authors").Find(&bks).Error != nil {
		return nil, errors.New("unable to find book records")
	}

	return bks, nil

}

// GetBook retrieve specific book record from the database.
func GetBook(r *http.Request) (Book, error) {

	bk := Book{}
	isbn := r.FormValue("isbn")

	if isbn == "" {
		return bk, errors.New("400. Bad Request")
	}

	aus := make([]Author, 0)
	config.Database.Model(&bk).Related(&aus, "Authors")
	if config.Database.Preload("Authors").Where(&Book{Isbn: isbn}).Find(&bk).Error != nil {
		return bk, errors.New("unable to find book for the given ISBN")
	}

	return bk, nil

}

// SaveBook create new book entry.
func SaveBook(r *http.Request) (Book, error) {

	// Get form values and validate
	bk, err := validateBookForm(r)

	if err != nil {
		return bk, err
	}

	tx := config.Database.Begin()

	if tx.Error != nil {
		return bk, errors.New("book cannot be saved")
	}

	if config.Database.Save(&bk).Error != nil {
		tx.Rollback()
		return bk, errors.New("book cannot be saved")
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return bk, errors.New("book cannot be saved")

	}

	return bk, nil

}

// UpdateBook modifies existing book details.
func UpdateBook(r *http.Request) (Book, error) {

	bk, err := validateBookForm(r)

	if err != nil {
		return bk, err
	}

	if err != nil {
		return bk, err
	}

	if err != nil {
		return bk, err
	}

	tx := config.Database.Begin()

	if tx.Error != nil {
		return bk, errors.New("book cannot be updated")
	}

	if config.Database.Where(&Book{Isbn: bk.Isbn}).Model(&bk).Updates(&bk).Error != nil {
		tx.Rollback()
		return bk, errors.New("unable to update book details")
	}

	// might not be the best way to update association
	// but calling just update does not really updates the associated records, it's just appending it
	if config.Database.Model(&bk).Association("Authors").Replace(bk.Authors).Error != nil {
		tx.Rollback()
		return bk, errors.New("unable to update book details")
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return bk, errors.New("unable to update book details")
	}

	return bk, nil

}

// DeleteBook removes book entry from the database.
func DeleteBook(r *http.Request) error {

	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	bk, err := GetBook(r)

	if err != nil {
		return errors.New("nothing to delete")
	}

	tx := config.Database.Begin()

	if tx.Error != nil {
		return errors.New("unable to delete book")
	}

	if config.Database.Model(&bk).Association("Authors").Delete(bk.Authors).Error != nil {
		tx.Rollback()
		return errors.New("unable to delete book")
	}

	if config.Database.Where(&Book{Isbn: isbn}).Delete(&Book{}).Error != nil {
		tx.Rollback()
		return errors.New("unable to delete book")
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return errors.New("unable to delete book")
	}

	return nil

}

// FormatBookPrice formats the price of the book.
func (bk *Book) FormatBookPrice() string {

	return fmt.Sprintf("$%.2f", bk.Price)

}

func validateBookForm(r *http.Request) (Book, error) {

	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	p := r.FormValue("price")
	a := r.Form["author"]

	if bk.Isbn == "" || bk.Title == "" {
		return bk, errors.New("fields cannot be empty")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("price must be a number")
	}

	bk.Price = float32(f64)

	alen := len(a)

	if alen <= 0 {
		return bk, errors.New("author ID can't be processed")
	}

	ids := make([]int32, alen)

	for i := 0; i < alen; i++ {
		int64, err := strconv.ParseInt(a[i], 10, 64)
		if err != nil {
			return bk, errors.New("author ID can't be processed")
		}
		ids[i] = int32(int64)
	}

	aus, err := GetByID(ids...)

	if err != nil {
		return bk, err
	}

	bk.Authors = aus

	return bk, nil

}

// ConvertToJSON converts the the book struct to json object
func (bk *Book) ConvertToJSON(w http.ResponseWriter) {

	uj, err := json.Marshal(bk)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

// GetAuthorNames concatenates all the book authors
func (bk *Book) GetAuthorNames() string {

	var b bytes.Buffer
	length := len(bk.Authors)
	for i := 0; i < length; i++ {
		b.WriteString(bk.Authors[i].Firstname)
		b.WriteString(" ")
		b.WriteString(bk.Authors[i].Lastname)
		if i < length-1 {
			b.WriteString(",")
		}
	}

	return b.String()
}

// GetAuthorIds place all author ids into a slice
func (bk *Book) GetAuthorIds() string {

	length := len(bk.Authors)
	ids := make([]int32, length)
	for i := 0; i < length; i++ {
		ids[i] = bk.Authors[i].AuthorID
	}

	uj, err := json.Marshal(ids)
	if err != nil {
		return ""
		fmt.Println(err)
	}

	return string(uj)
}

// GetAuthorDetails get all author details as separate paragraphs
func (bk *Book) GetAuthorDetails() []string {

	length := len(bk.Authors)
	details := make([]string, length)
	for i := 0; i < length; i++ {
		details[i] = bk.Authors[i].About
	}

	return details

}
