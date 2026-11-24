package websock

import (
	"github.com/kudryashov216/goWebSock/internal/wsframe"
)

func CurrentVersion() string {
	return CURRENTVERSION
}

func StartReader(ws *GoWebSock) {

	for {
		ws.read()
	}

}

/*External*/

func (ws *GoWebSock) SetConnect(headers string) error {

	_, err := ws.Conn.Write([]byte(headers))

	return err

}

func (ws *GoWebSock) Read(buff *[]byte) (n int) {

	return 0

}

func (ws *GoWebSock) Send(data []byte) (n int, err error) {

	frame := wsframe.NewFrame(&data)
	builder := wsframe.NewBuilder(frame)
	defer builder.Close()

	response, _ := builder.Build(0x01, &data)

	return ws.Conn.Write(*response)

}

func (ws *GoWebSock) SendToString(text string) (n int, err error) {

	buff := []byte(text)

	frame := wsframe.NewFrame(&buff)
	builder := wsframe.NewBuilder(frame)
	defer builder.Close()

	response, _ := builder.Build(0x01, &buff)

	return ws.Conn.Write(*response)

}

func (ws *GoWebSock) GetResponse() *[]byte {
	return &ws.Response
}

func (ws *GoWebSock) GetPayload() *[]byte {
	return &ws.Payload
}

func (ws *GoWebSock) GetPayloadToString() string {
	return string(ws.Payload)
}

func (ws *GoWebSock) SoftClose(text string) error {

	var buff []byte

	frame := wsframe.NewFrame(&buff)
	builder := wsframe.NewBuilder(frame)

	reasonClosure := []byte(text)

	buff = append(buff, 0x03, 0xE8)
	buff = append(buff, reasonClosure...)

	payload, _ := builder.Build(0x08, &buff)

	_, err := ws.Write(*payload)

	if err == nil {
		ws.Close()
	}

	return err

}

func (ws *GoWebSock) Close() error {

	return ws.Conn.Close()

}

/*Static*/

func (ws *GoWebSock) read() (n int, err error) {

	var buff []byte = make([]byte, 4096)

	frame := wsframe.NewFrame(&buff)
	reader := wsframe.NewReader(frame)
	defer reader.Close()

	code, n, err := reader.Read(ws.Conn, &buff)

	if code == 0x08 {
		ws.SoftClose("The connection was terminated at the client’s initiative.")
		ws.Close()
	}
	payload, _ := frame.Unmask(&buff)

	ws.Frames = append(ws.Frames, buff)
	ws.CurrentPayload = *payload
	ws.CurrentResponse = buff
	ws.NBytes = n

	return

}
