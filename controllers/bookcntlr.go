package controllers

import (
	"database/sql"
	"github.com/rburawes/golang-demo/config"
	"github.com/rburawes/golang-demo/models"
	"github.com/rburawes/golang-demo/sessions"
	"net/http"
)

// Index is the default page of the application.
func Index(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", &bks)

}

// Show provide book details.
func Show(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.GetBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "show.gohtml", &bk)

}

// CreateProcess validates the request method call the Create method
// to execute adding of new book.
func CreateProcess(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.SaveBook(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	bk.ConvertToJSON(w)

}

// Update modifies the book details.
func Update(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.GetBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk.ConvertToJSON(w)

}

// UpdateProcess validates request method and call Update method to
// modify book details.
func UpdateProcess(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	config.TPL.ExecuteTemplate(w, "success.gohtml", bk)

}

// DeleteProcess removes the book from the database.
func DeleteProcess(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := models.DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
