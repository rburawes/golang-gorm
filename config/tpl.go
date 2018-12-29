package config

import "html/template"

// TPL is the template object the will be
// available to the controllers for retrieving template files.
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
