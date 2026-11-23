package websock

type WebSock interface {
	SetConnect(headers string) error
	Send(data string) (n int, err error)
	Read(buff *[]byte) (n int, err error)
	Close() error
}
