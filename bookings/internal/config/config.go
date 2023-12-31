package config

import (
	"html/template"
	"github.com/alexedwards/scs/v2"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache		bool
	TemplateCache   map[string]*template.Template
	InfoLog			*log.Logger
	InProduction 	bool
	Session			*scs.SessionManager
}