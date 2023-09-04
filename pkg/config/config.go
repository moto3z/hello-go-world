package config

import "html/template"

// AppConfig holds th eapp config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
