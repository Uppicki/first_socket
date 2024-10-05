package wsservice

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type IWSService interface {
	ServeWS(
		http.ResponseWriter,
		*http.Request,
		http.Header,
	) error
	CreateConnection(
		http.ResponseWriter,
		*http.Request,
		http.Header,
	) (*websocket.Conn, error)
	Listen()
}
