package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// tc is a template cache, mapping the name to the template ptr
var tc = make(map[string]*template.Template)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if template exists in cache already
	_, inMap := tc[t]

	if !inMap {
		// Create template and add it
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		}
		log.Println("adding new template to cache")
	} else {
		log.Println("using existing template from cache")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)

}

// createTemplateCache updates the cache with a new template
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}
