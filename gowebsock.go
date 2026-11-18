package goWebSock

func CreateNewWebSocket(GET, Host, Upgrade, Connection, SecWebSocketKey, Origin string, SecWebSocketVersion uint8) GoWebSock {

	var ws GoWebSock

	ws.GET = GET
	ws.Host = Host
	ws.Origin = Origin
	ws.Upgrade = Upgrade
	ws.Connection = Connection
	ws.SecWebSocketKey = SecWebSocketKey
	ws.SecWebSocketVersion = SecWebSocketVersion

	return ws

}
