package main

import (
	"net/http"

	"github.com/gorilla/context"
	"lpp"
)

var frontend = "127.0.0.1:8001"
var backend = "ws://127.0.0.1:8002/chat"

func main() {

	lph := lpp.NewLongPollHandler(backend)

	http.ListenAndServe(frontend, context.ClearHandler(lph))
}
