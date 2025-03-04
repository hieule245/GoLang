package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/golang-web/internal/config"
	"github.com/golang-web/internal/models"
	"github.com/justinas/nosurf"
)

type mapCache map[string]*template.Template

// func RenderTemplate(w http.ResponseWriter, html string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.gohtml")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Fprintf(w, "Can't find html file")
// 	}
// }

var app *config.AppConfig

// NewTemplate to create a new template
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, gohtml string, data *models.TemplateData) {
	var tc mapCache

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[gohtml]
	if !ok {
		log.Fatal(ok)
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)

	err := t.Execute(buf, data)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (mapCache, error) {
	var myCache = make(mapCache)
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
