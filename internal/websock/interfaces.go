package websock

type WebSock interface {
	SetConnect() error
	Close() error
}
