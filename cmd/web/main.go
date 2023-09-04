package main

import (
	"fmt"
	"hello-world/pkg/config"
	"hello-world/pkg/handlers"
	"hello-world/pkg/render"
	"net/http"
)

const port = ":8080"

// main entry point of this app
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		print(err)
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting app : %s%s", "http://localhost", port)
	_ = http.ListenAndServe(port, nil)
}
