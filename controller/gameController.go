package controller

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lqd117/Game/config"
	"github.com/lqd117/Game/template"
	"net/http"
)

func GameHandler(w http.ResponseWriter, _ *http.Request) {
	template.RenderTemplate(w, "game", nil)
}

func WsGameHandler(w http.ResponseWriter, r *http.Request) {
	cnn, err := config.UpGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error", err)
		return
	}
	_ = cnn.WriteMessage(websocket.TextMessage, []byte("OK"))
}
