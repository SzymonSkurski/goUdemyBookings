package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/SzymonSkursrki/goUdemyBookings/pkg/config"
	"github.com/SzymonSkursrki/goUdemyBookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// add default data for all
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	//default data for all templates
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplatesCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from:", tmpl) //404 here
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td) //inject extra data here if not nill
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateTemplatesCache() (map[string]*template.Template, error) {
	path := "./templates/"
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob(path + "*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//check if match layout
		matches, err := filepath.Glob(path + "*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(path + "*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
