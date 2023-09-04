package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//Step1 : create a template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//Step2 : get request from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
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

// CreateTemplateCache is 캐싱
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
