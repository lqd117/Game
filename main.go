package main

import (
	"fmt"
	_ "github.com/lqd117/Game/memory"
	"github.com/lqd117/Game/session"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("template/index.html", "template/game.html"))
var globalSession *session.Manager

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

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "index", nil)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionId")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/index", 301)
		return
	}
	renderTemplate(w, "game", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	id, pwd := r.FormValue("id"), r.FormValue("pwd")
	flag := checkIdAndPassword(id, pwd)
	if flag == false {
		_, err := fmt.Fprintf(w, "没有该用户！")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	sess := globalSession.SessionStart(w, r)
	name := sess.Get(id)
	if name == nil {
		err := sess.Set(id, id)
		if err != nil {
			panic(err)
		}
	}
	http.Redirect(w, r, "/game", 301)
}

func init() {
	var err error
	globalSession, err = session.NewManager("memory", "sessionId", 3600)
	if err != nil {
		panic(err)
	}
	go globalSession.GC()
}

func main() {
	http.HandleFunc("/", makeHandler(indexHandler))
	http.HandleFunc("/index", makeHandler(indexHandler))
	http.HandleFunc("/login", makeHandler(loginHandler))
	http.HandleFunc("/game", makeHandler(gameHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {

	}
}
