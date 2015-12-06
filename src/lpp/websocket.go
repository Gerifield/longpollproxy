package lpp

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WebsocketServer struct {
	Ws *websocket.Conn
}

func NewWebsocketServer(backend string, header http.Header) (*WebsocketServer, error) {
	//delete(header, "Connection")
	ws, _, err := websocket.DefaultDialer.Dial(backend, nil)
	if err != nil {
		return nil, err
	}
	return &WebsocketServer{
		Ws: ws,
	}, nil
}

func (ws *WebsocketServer) Process() {

}

func (ws *WebsocketServer) SendSock(b []byte) error {
	return ws.Ws.WriteMessage(websocket.TextMessage, b)
}

func (ws *WebsocketServer) ReadSock() ([]byte, error) {
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

//TODO: Add channels etc
func (ws *WebsocketServer) Read() ([]byte, error) {
	_, b, err := ws.Ws.ReadMessage()
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}
