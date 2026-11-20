package websock

import (
	"net"
)

type GoWebSock struct {
	net.Conn
	Version string
}
