package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"hello-world/pkg/config"
	"hello-world/pkg/handlers"
	"hello-world/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main entry point of this app
func main() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	/*


	 */

	tc, err := render.CreateTemplateCache()
	if err != nil {
		println(err)
		log.Fatal("cannot create template cache")
	}

	// 1구간 ==========================
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplate(&app) //앱컨피그에 접근할 수 있도록
	//& 레퍼런스
	// * 포인터

	//2구간 ======================

	fmt.Printf("Starting app : %s%s", "http://localhost", portNumber)
	fmt.Println()

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
