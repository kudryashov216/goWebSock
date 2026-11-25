package goWebSock

import (
	"github.com/kudryashov216/goWebSock/internal"
	"github.com/kudryashov216/goWebSock/internal/websock"
	"github.com/kudryashov216/goWebSock/internal/wsframe"
)

type (
	RequestRW      internal.RequestRW
	WebSock        websock.WebSock
	ContextWebSock websock.ContextWebSock
)

type (
	WSFrame interface {
		wsframe.WSFrame
	}
)
