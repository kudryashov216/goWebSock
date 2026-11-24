package wsframe

import "net"

type WSFrame interface {
	CloserFrame
	Unmask(data *[]byte) (*[]byte, error)
}

type ReaderFrame interface {
	WSFrame
	Read(conn net.Conn, buff *[]byte) (code byte, n int, err error)
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
