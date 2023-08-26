package config

import (
	"html/template"
	"log"
)

// "Global" config file that can be imported throughout the project
// Avoid importing unnecessary and external packages in this kind of distributed file

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
