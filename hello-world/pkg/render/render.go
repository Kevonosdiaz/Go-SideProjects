package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create tmpl cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Fetch tmpl from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// Error checking execution of the template
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// Render the tmpl
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache generates a template cache
func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// First get all *.page.tmpl files from ./templates	(ending in .html for the time being)
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cache, err
	}

	// Range through all *.page.tmpl files
	for _, page := range pages {
		// Extract base name from full path
		name := filepath.Base(page)
		// Parse it
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		// Look for layout files from ./templates matching the current page
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		// In the case of a matching layout, parse the layout and add it to (t)emplate (s)et
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		// Add the template set to the page name
		cache[name] = ts
	}

	return cache, nil
}

// // tc is a template cache, mapping the name to the template ptr
// var tc = make(map[string]*template.Template)
// // RenderTemplate renders a template
// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check if template exists in cache already
// 	_, inMap := tc[t]

// 	if !inMap {
// 		// Create template and add it
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		log.Println("adding new template to cache")
// 	} else {
// 		log.Println("using existing template from cache")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)

// }

// // createTemplateCache updates the cache with a new template
// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }
