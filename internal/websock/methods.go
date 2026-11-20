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

	var (
		header, payload []byte
	)

	payload = []byte(data)

	header = append(header, 0x81)

	switch {

	case len(payload) <= 125:

		header = append(header, byte(len(payload)))

	case len(payload) <= 65535:

		header = append(header, 126, byte(len(payload)>>8), byte(len(payload)))

	default:

		length := uint64(len(payload))

		header = append(header, 127)

		for i := 7; i >= 0; i-- {
			header = append(header, byte(length>>(8*i)))
		}

	}

	frame := append(header, payload...)

	return ws.Conn.Write(frame)

}

func (ws *GoWebSock) Close() error {

	return ws.Conn.Close()

}
