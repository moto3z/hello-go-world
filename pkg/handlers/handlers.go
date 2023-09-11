package handlers

import (
	"hello-world/pkg/config"
	"hello-world/pkg/models"
	"hello-world/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo 새로운 리포지토리 생성
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the Repos for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

//td *models.TemplateData

// Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	//ip 세션
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello world!"

	//ip세션
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	//m.App.Session.
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
