package render

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	pt, _ := template.ParseFiles("./templates/" + tmpl)
	err := pt.Execute(w, nil)
	if err != nil {
		print("error")
	}
}
