package websocket

import "github.com/kudryashov216/goWebSock/internal/enums"

type GoWebSock struct {
	SecWebSocketVersion uint8
	GET                 string
	Host                string
	Connection          string
	SecWebSocketKey     string
	Origin              string
	Upgrade             enums.TypeConn
}
