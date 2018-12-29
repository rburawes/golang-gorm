package routes

import (
	"github.com/rburawes/golang-demo/controllers"
	"net/http"
)

// LoadRoutes handles routes to pages of the application.
func LoadRoutes() {

	// Index or main page.
	http.HandleFunc("/", index)
	// Book related routes
	http.HandleFunc("/books", controllers.Index)
	http.HandleFunc("/books/show", controllers.Show)
	http.HandleFunc("/books/create/process", controllers.CreateProcess)
	http.HandleFunc("/books/update", controllers.Update)
	http.HandleFunc("/books/update/process", controllers.UpdateProcess)
	http.HandleFunc("/books/delete/process", controllers.DeleteProcess)
	// Author related route(s)
	http.HandleFunc("/authors", controllers.GetAuthors)
	// User related route(s)
	http.HandleFunc("/signup", controllers.Signup)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	// CSS, JS and images
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./resource"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Listens and serve requests.
	http.ListenAndServe(":8080", nil)

}

// Redirect to list of books.
func index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
