package goWebSock

import (
	"net/http"
)

type RequestRW struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}
