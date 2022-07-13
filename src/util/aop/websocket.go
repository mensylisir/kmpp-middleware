package aop

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var UpGrader websocket.Upgrader

func init() {
	UpGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}
