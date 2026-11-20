package websock

type WebSock interface {
	SetConnect() error
	Send(data string) (int, error)
	Close() error
}
