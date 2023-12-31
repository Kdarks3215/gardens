package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/Kdarks3215/gardens/pkg/config"
	"github.com/Kdarks3215/gardens/pkg/models"
)
var functions = template.FuncMap{
	 
}
var app *config.AppConfig
//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}


//RenderTemplate using html template
func RenderTemplate(w http.ResponseWriter, tmpl string,td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
	//get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buff := new(bytes.Buffer)

	td =  AddDefaultData(td)

	_ = t.Execute(buff, td)

	_,err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
	 
}
//CreateTemplateCache creates template cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache:= map[string]*template.Template{}


	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache,err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache,err
		}

		if len(matches) >0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache,err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}