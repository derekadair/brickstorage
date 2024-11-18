package views

import (
	"html/template"
	"brickstorage/domain/part"
)

func NewPartView(templ *template.Template) *ModelView[part.Part] {
	return NewModelView[part.Part](templ, "part")
}
