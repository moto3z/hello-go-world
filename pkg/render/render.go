package render

import (
	"bytes"
	"fmt"
	"hello-world/pkg/config"
	"hello-world/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app = &config.AppConfig{}

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// AddDefaultData 모름
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//Step1 : create a template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func 과거의코드(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		print("error")
	}
}

// CreateTemplateCache is 캐함
func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	// get all of the files named "*.page.tmpl"
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//rage theough all fils ending wilt
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// @todo Glob >> 아 HTML 경로와 이름 가지고
			// 에러가 좀 국한적이다 문제가 있다 >>
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}
