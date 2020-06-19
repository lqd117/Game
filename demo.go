package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID, pwd string
}

var templates = template.Must(template.ParseFiles("template/index.html", "template/game.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	id, pwd := r.FormValue("id"), r.FormValue("pwd")
	flag := checkIdAndPassword(id, pwd)
	if flag == false {
		fmt.Fprintf(w, "没有该用户！")
		return
	}
	renderTemplate(w, "game", &User{ID: id, pwd: pwd})
}

func main() {
	http.HandleFunc("/", makeHandler(indexHandler))
	http.HandleFunc("/index", makeHandler(indexHandler))
	http.HandleFunc("/login", makeHandler(loginHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {

	}
}
