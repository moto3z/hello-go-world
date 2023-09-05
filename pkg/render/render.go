package render

import (
	"bytes"
	"hello-world/pkg/config"
	"hello-world/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

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
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//Step1 : create a template cache
	if err != nil {
		log.Fatal(err)
	}

	//Step2 : get request from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("쿠드 낫 겟 쳄플릿 프롬 템캐시")
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err = t.Execute(buf, nil)
	_ = t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	//Step3 : render the tmpl
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	//before code
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
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

//var tc = make(map[string]*template.Template)
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//	//check see if we already have the template in our cache
//	_, inMap := tc[t]
//	if !inMap {
//		log.Println("creating template and adding to cache")
//		//need create this template
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		//have template in cache already
//		log.Println("using cache template")
//	}
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//}

//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintln("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//
//	tc[t] = tmpl
//	return nil
//}
