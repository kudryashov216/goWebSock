package creater

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/kudryashov216/goWebSock/internal"
	"github.com/kudryashov216/goWebSock/internal/data"
	"github.com/kudryashov216/goWebSock/internal/websock"
	"github.com/kudryashov216/goWebSock/internal/websockclient"
)

func NewRequestRW(rw http.ResponseWriter, r *http.Request) *internal.RequestRW {

	return &internal.RequestRW{ResponseWriter: rw, Request: r}

}

func CreateWebSockServerFromRequestRW(reqRW *internal.RequestRW) (websock.ContextWebSock, websock.WebSock, error) {

	hj, ok := reqRW.ResponseWriter.(http.Hijacker)

	if !ok {
		return nil, nil, errors.New("Failed to create a websocket.")
	}

	conn, _, err := hj.Hijack()

	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	headersMap := reqRW.Request.Header

	headersMap["GET"] = []string{reqRW.Request.Proto}
	headersMap["Host"] = []string{reqRW.Request.Host}

	headers := data.HeaderToString(&headersMap)

	ctx, ws, err := CreateWebSockServer(conn, headers)

	return ctx, ws, err

}

func CreateWebSockServer(conn net.Conn, headers string) (websock.ContextWebSock, websock.WebSock, error) {

	var (
		err           error
		parentContext context.Context = context.Background()
	)

	if len(headers) == 0 {

		var buff []byte = make([]byte, 4096)

		n, err := conn.Read(buff)

		if err != nil {
			conn.Close()
			return nil, nil, err
		}

		headers = string(buff[:n])

	}

	ok := data.HeadersIsValid(headers)

	if !ok {
		conn.Close()
		return nil, nil, errors.New("The headers are not valid for creating a websocket.")
	}

	headerMap := data.GetHeadersToMap(headers)
	headerMap = data.KeyMapToLower(headerMap)

	hash := sha1.New()
	hash.Write([]byte((*headerMap)["sec-websocket-key"] + websock.WSGUID))

	accept := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	responseHeaders := fmt.Sprintf("%s 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", (*headerMap)["get"], accept)

	ctx, cancel := context.WithCancel(parentContext)

	webSock := websock.GoWebSock{
		Conn:    conn,
		Version: websock.CURRENTVERSION,
		Ctx:     ctx,
		Сancel:  cancel,
	}

	err = webSock.SetConnect(responseHeaders)

	return &webSock.ContextWS, &webSock, err

}

func CreateWebSockClient(address string) (websockclient.WebSockClient, error) {

	_, err := net.Dial("tcp", address)

	return nil, err

}
