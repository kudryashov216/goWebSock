package goWebSock

import (
	"errors"
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
		return nil, errors.New("The headers are not valid for creating a websocket.")
	}

	response := ""

	webSock := websock.GoWebSock{
		Conn:     conn,
		Version:  websock.CURRENTVERSION,
		Response: response,
	}

	return webSock, nil

}
