package websock

func (ws *GoWebSock) SetConnect() error {

	_, err := ws.Conn.Write([]byte(ws.Response))
	return err

}

func (ws *GoWebSock) Close() error {
	return ws.Conn.Close()
}

func CurrentVersion() string {
	return CURRENTVERSION
}
