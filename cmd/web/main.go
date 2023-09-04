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

	// 1구간 ==========================
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplate(&app)
	//& 레퍼런스
	// * 포인터

	//2구간 ======================
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting app : %s%s", "http://localhost", port)
	_ = http.ListenAndServe(port, nil)
}
