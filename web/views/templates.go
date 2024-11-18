package views

import (
	"embed"
	"html/template"
)

var (
	//go:embed "templates/*"
	partTemplates embed.FS
)

func NewTemplates() (*template.Template, error) {
	return template.ParseFS(partTemplates, "templates/*/*.gohtml", "templates/*.gohtml")
}
