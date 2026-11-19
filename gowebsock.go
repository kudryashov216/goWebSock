package goWebSock

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
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

	headerMap, ok := data.HeadersIsValid(headers)

	if !ok {
		return nil, errors.New("The headers are not valid for creating a websocket.")
	}

	hash := sha1.New()
	hash.Write([]byte((*headerMap)["Sec-WebSocket-Key"] + websock.WSGUID))

	accept := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	response := fmt.Sprintf("%s 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n", (*headerMap)["Get"], accept)

	webSock := websock.GoWebSock{
		Conn:     conn,
		Version:  websock.CURRENTVERSION,
		Response: response,
	}

	webSock.SetConnect()

	return &webSock, nil

}
