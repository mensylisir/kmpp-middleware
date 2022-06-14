package aop

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
)

func WebsocketHandler(c *gin.Context) (*websocket.Conn, error) {
	ReadBufferSize := viper.GetInt("bind.read_buffer_size")
	WriteBufferSize := viper.GetInt("bind.write_buffer_size")
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriteBufferSize,
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			panic(err)
		}
	}()
	return ws, nil
}
