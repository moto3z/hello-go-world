package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the app config 전체 앱에서 공유하는거라서 퍼블릭으로 열어줘야함
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
