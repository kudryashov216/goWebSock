package internal

import (
	"net/http"
)

type RequestRW struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}
