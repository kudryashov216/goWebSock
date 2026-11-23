package websockclient

type WebSockClient interface {
	Read()
	Send()
	Close()
}
