package websock

import (
	"context"
	"net"

	"github.com/kudryashov216/goWebSock/internal/wsframe"
)

type WebSock interface {
	LocalAddress() net.Addr
	RemoteAddress() net.Addr
	ReadFrames() (wsframe.WSFrame, error)
	SetConnect(headers string) error
	Send(opcode byte, data []byte) (n int, err error)
	SendToString(text string) (n int, err error)
	SoftClose(text string) error
	Close() error
}

type ContextWebSock interface {
	GetContext() context.Context
	Cancel()
}
