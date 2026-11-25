package websock

import (
	"context"

	"github.com/kudryashov216/goWebSock/internal/wsframe"
)

func CurrentVersion() string {
	return CURRENTVERSION
}

/*External*/

/*WebSock*/
func (ws *GoWebSock) ReadFrames() (wsframe.WSFrame, error) {

	n, err := ws.read()

	if n == 0 {
		return nil, err
	}

	return ws.Frame, nil

}

func (ws *GoWebSock) SetConnect(headers string) error {

	_, err := ws.Conn.Write([]byte(headers))
	return err

}

func (ws *GoWebSock) Send(data []byte) (n int, err error) {

	frame := wsframe.NewFrame(ws.Ctx, &data)
	builder := wsframe.NewBuilder(frame)
	defer builder.Close()

	response, _ := builder.Build(0x01, &data)

	return ws.Conn.Write(*response)

}

func (ws *GoWebSock) SendToString(text string) (n int, err error) {

	buff := []byte(text)

	frame := wsframe.NewFrame(ws.Ctx, &buff)
	builder := wsframe.NewBuilder(frame)
	defer builder.Close()

	response, _ := builder.Build(0x01, &buff)

	return ws.Conn.Write(*response)

}

func (ws *GoWebSock) SoftClose(text string) error {

	var buff []byte

	frame := wsframe.NewFrame(ws.Ctx, &buff)
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

/*ContextWS*/
func (cws *ContextWS) GetContext() context.Context {
	return cws.Ctx
}

func (cws *ContextWS) Cancel() {
	cws.Сancel()
}

/*Static*/

func (ws *GoWebSock) read() (n int, err error) {

	var buff = make([]byte, 4096)

	frame := wsframe.NewFrame(ws.Ctx, &buff)
	reader := wsframe.NewReader(frame)
	defer reader.Close()

	code, n, err := reader.Read(ws.Conn, &buff)

	if code == 0x08 {
		text := "The connection was terminated at the client’s initiative."
		ws.SoftClose(text)
		ws.ContextWS.Сancel()
		ws.Close()
	}

	ws.Frame = frame

	return

}
