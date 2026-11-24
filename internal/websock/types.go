package websock

import (
	"net"
)

type GoWebSock struct {
	net.Conn
	Version         string
	Frames          [][]byte
	CurrentResponse []byte
	CurrentPayload  []byte
	NBytes          int
}
