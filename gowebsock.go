package goWebSock

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/kudryashov216/goWebSock/internal/data"
	"github.com/kudryashov216/goWebSock/internal/websock"
)

func NewRequestRW(rw http.ResponseWriter, r *http.Request) *RequestRW {

	return &RequestRW{ResponseWriter: rw, Request: r}

}

func CreateWebSockServerFromRequestRW(reqRW *RequestRW) (websock.WebSock, error) {

	hj, ok := reqRW.ResponseWriter.(http.Hijacker)

	if !ok {
		return nil, errors.New("Failed to create a websocket.")
	}

	conn, _, err := hj.Hijack()

	if err != nil {
		conn.Close()
		return nil, err
	}

	headersMap := reqRW.Request.Header

	headersMap["GET"] = []string{reqRW.Request.Proto}
	headersMap["Host"] = []string{reqRW.Request.Host}

	headers := data.HeaderToString(&headersMap)

	ws, err := CreateWebSockServer(conn, headers)

	return ws, err

}

func CreateWebSockServer(conn net.Conn, headers string) (websock.WebSock, error) {

	var err error

	if len(headers) == 0 {

		var buff []byte = make([]byte, 4096)

		n, err := conn.Read(buff)

		if err != nil {
			conn.Close()
			return nil, err
		}

		headers = string(buff[:n])

	}

	ok := data.HeadersIsValid(headers)

	if !ok {
		conn.Close()
		return nil, errors.New("The headers are not valid for creating a websocket.")
	}

	headerMap := data.GetHeadersToMap(headers)
	headerMap = data.KeyMapToLower(headerMap)

	hash := sha1.New()
	hash.Write([]byte((*headerMap)["sec-websocket-key"] + websock.WSGUID))

	accept := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	responseHeaders := fmt.Sprintf("%s 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", (*headerMap)["get"], accept)

	webSock := websock.GoWebSock{
		Conn:    conn,
		Version: websock.CURRENTVERSION,
	}

	err = webSock.SetConnect(responseHeaders)

	return &webSock, err

}

func CreateWebSockClient(address string) (websock.WebSock, error) {

	_, err := net.Dial("tcp", address)

	return nil, err

}
