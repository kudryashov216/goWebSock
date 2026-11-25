package goWebSock

import (
	"net"
	"net/http"

	"github.com/kudryashov216/goWebSock/internal"
	"github.com/kudryashov216/goWebSock/internal/creater"
)

func NewRequestRW(rw http.ResponseWriter, r *http.Request) *RequestRW {

	reqRW := creater.NewRequestRW(rw, r)

	return (*RequestRW)(reqRW)

}

func CreateWebSockServerFromRequestRW(reqRW *RequestRW) (ContextWebSock, WebSock, error) {

	return creater.CreateWebSockServerFromRequestRW((*internal.RequestRW)(reqRW))

}

func CreateWebSockServer(conn net.Conn, headers string) (ContextWebSock, WebSock, error) {
	return creater.CreateWebSockServer(conn, headers)
}
