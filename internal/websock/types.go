package websock

import (
	"context"
	"net"

	"github.com/kudryashov216/goWebSock/internal/wsframe"
)

type GoWebSock struct {
	net.Conn
	ContextWS
	Version string
	Frame   wsframe.WSFrame
}

type ContextWS struct {
	Ctx    context.Context
	Сancel context.CancelFunc
}
