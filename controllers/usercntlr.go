package controllers

import (
	"github.com/rburawes/golang-demo/config"
	"github.com/rburawes/golang-demo/models"
	"github.com/rburawes/golang-demo/sessions"
	"net/http"
	"time"
)

// Signup allows the user to create an account.
func Signup(w http.ResponseWriter, r *http.Request) {

	if sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if r.Method == http.MethodPost {

		u, err := models.SaveUser(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// create session
		sessions.CreateSession(w, u)

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", u)

}

// Login allows registered user to access the application.
func Login(w http.ResponseWriter, r *http.Request) {

	if sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		e := r.FormValue("email")

		if un == "" {
			un = e
		}

		// check if the user exists
		u, ok := models.FindUser(un)
		if !ok {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		if !u.ValidatePassword(p) {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		sessions.CreateSession(w, u)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	config.TPL.ExecuteTemplate(w, "login.gohtml", u)

}

// Logout method to call when the user signed out of the application.
func Logout(w http.ResponseWriter, r *http.Request) {

	if !sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sessions.RemoveSession(w, r)

	//not the best place and not to be used in production
	if time.Now().Sub(sessions.StoredSessionClean) > (time.Second * 30) {
		go sessions.CleanSessions()
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
