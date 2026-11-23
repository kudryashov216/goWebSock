package websockclient

import (
	"net"
)

type GoWebSockClient struct {
	net.Conn
	Version string
}
