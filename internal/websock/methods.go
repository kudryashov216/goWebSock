package websock

func CurrentVersion() string {
	return CURRENTVERSION
}

func (ws *GoWebSock) SetConnect() error {

	_, err := ws.Conn.Write([]byte(ws.Response))

	return err

}

func (ws *GoWebSock) Send(data string) (int, error) {

	n, err := ws.Write([]byte(data))

	return n, err

}

func (ws *GoWebSock) Close() error {

	return ws.Conn.Close()

}
