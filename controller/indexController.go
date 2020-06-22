package controller

import (
	"fmt"
	"github.com/lqd117/Game/config"
	"github.com/lqd117/Game/dataBase"
	"github.com/lqd117/Game/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	template.RenderTemplate(w, "index", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	id, pwd := r.FormValue("id"), r.FormValue("pwd")
	flag := dataBase.CheckIdAndPassword(id, pwd)
	if flag == false {
		_, err := fmt.Fprintf(w, "没有该用户！")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	sess := config.GlobalSession.SessionStart(w, r)
	name := sess.Get(id)
	if name == nil {
		err := sess.Set(id, id)
		if err != nil {
			panic(err)
		}
	}
	http.Redirect(w, r, "/game", 301)
}
