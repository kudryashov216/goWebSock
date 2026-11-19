package websock

func (ws *GoWebSock) SetConnect() {

	ws.Conn.Write([]byte(ws.Response))

}

func (ws *GoWebSock) Close() {

}

func CurrentVersion() string {
	return CURRENTVERSION
}
