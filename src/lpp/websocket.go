package lpp

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WebsocketServer struct {
	Ws *websocket.Conn

	//Read buffer
	ReadCh      chan string
	ReaderError error
}

type ReadTimeoutError struct{}

func (ReadTimeoutError) Error() string {
	return "Read timeout"
}

func NewWebsocketServer(backend string, header http.Header) (*WebsocketServer, error) {
	//delete(header, "Connection")
	ws, _, err := websocket.DefaultDialer.Dial(backend, nil)
	if err != nil {
		return nil, err
	}
	return &WebsocketServer{
		Ws:     ws,
		ReadCh: make(chan string, 256),
	}, nil
}

func (ws *WebsocketServer) ProcessRead() {
	for {
		b, err := ws.readSock()
		if err != nil {
			ws.ReaderError = err
		}
		ws.ReadCh <- string(b)
	}
}

func (ws *WebsocketServer) readSock() ([]byte, error) {
	_, b, err := ws.Ws.ReadMessage()
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

//TODO: Add channels etc
func (ws *WebsocketServer) Send(b []byte) error {
	return ws.Ws.WriteMessage(websocket.TextMessage, b)
}

func (ws *WebsocketServer) Read() ([]byte, error) {
	if ws.ReaderError != nil {
		return []byte(""), ws.ReaderError
	}

	select {
	case str := <-ws.ReadCh:
		return []byte(str), nil
	case <-time.After(5 * time.Second):
		return []byte(""), ReadTimeoutError{}
	}

}
