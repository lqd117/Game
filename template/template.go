package template

import (
	"github.com/lqd117/Game/config"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := config.Templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
