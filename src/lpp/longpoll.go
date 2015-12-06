package lpp

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/sessions"
	//"github.com/gorilla/websocket"
)

type LongPollServer struct {
	Backend string
	Session *sessions.CookieStore
	//ConnectionStore map[string]*websocket.Conn
	ConnectionStore map[interface{}]*WebsocketServer
}

func NewLongPollHandler(backend string) *LongPollServer {
	return &LongPollServer{
		Backend:         backend,
		Session:         sessions.NewCookieStore([]byte("enc-key-for-session")),
		ConnectionStore: make(map[interface{}]*WebsocketServer),
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
	socketId, ok := session.Values["socketId"]
	if !ok {
		socketId = genRandomId()
		session.Values["socketId"] = socketId
	}
	session.Save(r, w)

	log.Println(r.Header)

	var ws *WebsocketServer
	ws, ok = lps.ConnectionStore[socketId]
	if !ok {
		ws, err = NewWebsocketServer(lps.Backend, r.Header)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		lps.ConnectionStore[socketId] = ws
	}

	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		log.Println(string(b))
		err = ws.Send(b)
		if err != nil {
			http.Error(w, err.Error(), 201)
			return
		}
	} else {
		//GET
		b, err := ws.Read()
		if err != nil {
			http.Error(w, err.Error(), 201)
			return
		}
		w.Write(b)
	}
}

func genRandomId() string {
	return fmt.Sprintf("%d%d", rand.Int63(), rand.Int63())
}
