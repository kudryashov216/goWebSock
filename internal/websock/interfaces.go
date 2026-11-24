package websock

type WebSock interface {
	SetConnect(headers string) error
	Read(buff *[]byte) (n int)
	Send(data []byte) (n int, err error)
	SendToString(text string) (n int, err error)
	GetResponse() *[]byte
	GetPayload() *[]byte
	GetPayloadToString() string
	SoftClose(text string) error
	Close() error
}
