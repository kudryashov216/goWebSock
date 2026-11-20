package websock

type WebSock interface {
	SetConnect(headers string) error
	Send(data string) (int, error)
	Close() error
}
