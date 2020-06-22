package main

import (
	"github.com/lqd117/Game/config"
	"github.com/lqd117/Game/controller"
	"github.com/lqd117/Game/filter"
	"github.com/lqd117/Game/session"
	"net/http"
)

func init() {
	var err error
	config.GlobalSession, err = session.NewManager("memory", "sessionId", 3600)
	if err != nil {
		panic(err)
	}
	go config.GlobalSession.GC()
}

func main() {
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/index", controller.IndexHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/game", filter.HttpFilterHandler(controller.GameHandler))
	http.HandleFunc("/ws/game", filter.WsFilterHandler(controller.WsGameHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {

	}
}
