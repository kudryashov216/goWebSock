package websock

import (
	"context"
	"io"
	"net"

	"github.com/kudryashov216/goWebSock/internal/wsframe"
	"github.com/kudryashov216/goWebSock/opcodes"
)

func CurrentVersion() string {
	return CURRENTVERSION
}

/*External*/

/*WebSock*/
func (ws *GoWebSock) ReadFrames() (wsframe.WSFrame, error) {

	n, err := ws.read()
	ws.Frame.(*wsframe.Frame).NBytes = n

	return ws.Frame, err

}

func (ws *GoWebSock) SetConnect(headers string) error {

	_, err := ws.Conn.Write([]byte(headers))
	return err

}

func (ws *GoWebSock) Send(opcode byte, data []byte) (n int, err error) {

	frame := wsframe.NewFrame(ws.Ctx, &data)
	builder := wsframe.NewBuilder(frame)
	defer builder.Close()

	response, _ := builder.Build(opcode, &data)

	return ws.Conn.Write(*response)

}

func (ws *GoWebSock) SendToString(text string) (n int, err error) {

	buff := []byte(text)

	frame := wsframe.NewFrame(ws.Ctx, &buff)
	builder := wsframe.NewBuilder(frame)
	defer builder.Close()

	response, _ := builder.Build(opcodes.STRING, &buff)

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

func (ws *GoWebSock) RemoteAddress() net.Addr {

	return ws.Conn.RemoteAddr()

}

func (ws *GoWebSock) LocalAddress() net.Addr {

	return ws.Conn.LocalAddr()

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
		err = io.EOF
		text := "The connection was terminated at the client’s initiative."
		ws.SoftClose(text)
		ws.ContextWS.Сancel()
		ws.Close()
	}

	ws.Frame = frame

	return

}
