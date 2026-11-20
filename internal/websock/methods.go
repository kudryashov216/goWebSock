package websock

func CurrentVersion() string {
	return CURRENTVERSION
}

func (ws *GoWebSock) SetConnect(headers string) error {

	_, err := ws.Conn.Write([]byte(headers))

	return err

}

func (ws *GoWebSock) Read(buff []byte) (n int, err error) {

	return ws.Conn.Read(buff)

}

func (ws *GoWebSock) Send(data string) (int, error) {

	return ws.Conn.Write([]byte(data))

}

func (ws *GoWebSock) Close() error {

	return ws.Conn.Close()

}
