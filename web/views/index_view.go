package views

import (
	"html/template"
	"brickstorage/domain/part"
	"net/http"
)

type IndexView struct {
	templ *template.Template
}

func NewIndexView(templ *template.Template) *IndexView {
	return &IndexView{templ: templ}
}

func (t *IndexView) Index(w http.ResponseWriter, parts []part.Part) {
	var viewModel any = parts
	if err := t.templ.ExecuteTemplate(w, "index", viewModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
