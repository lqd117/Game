package config

import (
	"github.com/gorilla/websocket"
	"github.com/lqd117/Game/session"
	"html/template"
	"net/http"
)

var (
	Templates     = template.Must(template.ParseFiles("template/index.html", "template/game.html"))
	GlobalSession *session.Manager
	UpGrader      = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
