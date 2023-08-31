package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateV0(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		print("error")
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	//check see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache")
		//need create this template
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//have template in cache already
		log.Println("using cache template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintln("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
