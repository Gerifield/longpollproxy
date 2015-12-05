package lpp

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type LongPollServer struct {
	Backend string
	Session *sessions.CookieStore
}

func NewLongPollHandler(backend string) *LongPollServer {
	return &LongPollServer{
		Backend: backend,
		Session: sessions.NewCookieStore([]byte("enc-key-for-session")),
	}
}

func (lps *LongPollServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}
	session, err := lps.Session.Get(r, "sessid")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	session.Save(r, w)
	log.Println(session.ID)
	log.Println(session.IsNew)
	log.Println("OK?")
}
