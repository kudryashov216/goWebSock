package wsframe

import "net"

type WSFrame interface {
	CloserFrame
}

type ReaderFrame interface {
	WSFrame
	Read(conn net.Conn, buff *[]byte) (n int, err error)
}

type BuilderFrame interface {
	WSFrame
	Build(opcode byte, buff *[]byte) (data *[]byte, err error)
}

type RBFrame interface {
	ReaderFrame
	BuilderFrame
}

type CloserFrame interface {
	Close()
	IsClosed() bool
}
