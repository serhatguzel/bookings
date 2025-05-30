package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/serhatguzel/bookings/internal/config"
	"github.com/serhatguzel/bookings/internal/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	td.Flash = app.Session.PopString(r.Context(), "flash")     // get the flash message from the session
	td.Error = app.Session.PopString(r.Context(), "error")     // get the error message from the session
	td.Warning = app.Session.PopString(r.Context(), "warning") // get the warning message from the session

	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders a template using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		//Diskten erişmiyoruz artık
		tc = app.TemplateCache // get the template cache from the app config

	} else {

		tc, _ = CreateTemplateCache()

	}

	// get the template from the cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r) // add default data to the template data

	_ = t.Execute(buf, td)

	// render the template to the response writer

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from the ./templates/ directory
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through the pages and create a template cache
	for _, page := range pages {
		name := filepath.Base(page)                    // get the base name of the file (e.g. home.page.tmpl)
		ts, err := template.New(name).ParseFiles(page) // parse the page file
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl") // get all the layout files
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") // parse the layout files
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts // add the template to the cache
	}

	return myCache, nil
}
