package websock

import (
	"context"

	"github.com/kudryashov216/goWebSock/internal/wsframe"
)

type WebSock interface {
	ReadFrames() (wsframe.WSFrame, error)
	SetConnect(headers string) error
	Send(data []byte) (n int, err error)
	SendToString(text string) (n int, err error)
	SoftClose(text string) error
	Close() error
}

type ContextWebSock interface {
	GetContext() context.Context
	Cancel()
}
