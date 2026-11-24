package goWebSock

import (
	"net"
	"net/http"

	"github.com/kudryashov216/goWebSock/internal"
	"github.com/kudryashov216/goWebSock/internal/creater"
	"github.com/kudryashov216/goWebSock/internal/websock"
)

func NewRequestRW(rw http.ResponseWriter, r *http.Request) *internal.RequestRW {

	return creater.NewRequestRW(rw, r)

}

func CreateWebSockServerFromRequestRW(reqRW *internal.RequestRW) (websock.WebSock, error) {

	return creater.CreateWebSockServerFromRequestRW(reqRW)

}

func CreateWebSockServer(conn net.Conn, headers string) (websock.WebSock, error) {
	return creater.CreateWebSockServer(conn, headers)
}
