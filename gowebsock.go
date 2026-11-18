package goWebSock

import (
	"github.com/kudryashov216/goWebSock/internal/enums"
	"github.com/kudryashov216/goWebSock/internal/websocket"
)

func CreateNewWebSocket(GET, Host, Connection, SecWebSocketKey, Origin string, Upgrade enums.TypeConn, SecWebSocketVersion uint8) websocket.WebSock {

	var ws websocket.GoWebSock

	ws.GET = GET
	ws.Host = Host
	ws.Origin = Origin
	ws.Upgrade = Upgrade
	ws.Connection = Connection
	ws.SecWebSocketKey = SecWebSocketKey
	ws.SecWebSocketVersion = SecWebSocketVersion

	return ws

}
