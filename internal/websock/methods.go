package websock

import "github.com/kudryashov216/goWebSock/internal/wsframe"

func CurrentVersion() string {
	return CURRENTVERSION
}

func (ws *GoWebSock) SetConnect(headers string) error {

	_, err := ws.Conn.Write([]byte(headers))

	return err

}

func (ws *GoWebSock) Read(buff *[]byte) (n int, err error) {

	frame := wsframe.NewFrame(buff)
	reader := wsframe.NewReader(frame)
	defer reader.Close()

	code, n, err := reader.Read(ws.Conn, buff)

	if code == 0xFF {
		ws.Close()
		return
	}

	return

}

func (ws *GoWebSock) Send(data string) (n int, err error) {

	var buff []byte

	buff = []byte(data)

	frame := wsframe.NewFrame(&buff)
	builder := wsframe.NewBuilder(frame)
	defer builder.Close()

	response, _ := builder.Build(0x01, &buff)

	return ws.Conn.Write(*response)

}

func (ws *GoWebSock) Close() error {

	return ws.Conn.Close()

}
