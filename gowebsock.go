package goWebSock

import (
	"net"

	"github.com/kudryashov216/goWebSock/internal/data"
	"github.com/kudryashov216/goWebSock/internal/websock"
)

func CreateWebSockServer(conn net.Conn) (websock.WebSock, error) {

	buff := make([]byte, 4096)

	n, err := conn.Read(buff)

	if err != nil {
		return nil, err
	}

	headers := string(buff[:n])

	if !data.HeadersIsValid(headers) {
		return nil, nil
	}

	return nil, nil

}
