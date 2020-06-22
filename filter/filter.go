package filter

import (
	"github.com/lqd117/Game/config"
	"net/http"
)

func HttpFilterHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionId")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/index", 301)
			return
		}
		fn(w, r)
	}
}

func WsFilterHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sid := r.URL.Query().Get("sessionId")
		sess := config.GlobalSession.SessionFinder(sid)
		if sess == nil {
			return
		}
		fn(w, r)
	}
}
