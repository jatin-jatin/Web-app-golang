package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	// render the template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(fmt.Errorf("requested file/template missing for %s", tmpl))
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println("error parsing template", err)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error parsing template", err)
		return
	}
	// parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	// err := parsedTemplate.Execute(w, nil)
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//  get all of the files named *page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}
