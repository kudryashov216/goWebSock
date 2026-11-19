package websock

type WebSock interface {
	Connect()
	CurrentVersion() uint8
}
