package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rburawes/golang-demo/models"
	"github.com/rburawes/golang-demo/sessions"
	"net/http"
)

// GetAuthors returns all the available authors.
func GetAuthors(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden) // 403
		fmt.Fprintf(w, "%s\n", "Access to this resource was denied!")
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	aus, err := models.AllAuthors()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(aus)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)

}
