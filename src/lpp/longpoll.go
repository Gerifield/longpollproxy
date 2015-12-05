package lpp

import (
	"log"
	"net/http"
)

type LongPollServer struct {
	Backend string
}

func NewLongPollHandler(backend string) *LongPollServer {
	return &LongPollServer{
		Backend: backend,
	}
}

func (lps *LongPollServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("GET", r.RequestURI)
	case "POST":
		log.Println("POST", r.RequestURI)
	default:
		http.Error(w, "Wrong method", http.StatusBadRequest)
	}
}
