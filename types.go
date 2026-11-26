package goWebSock

import (
	"github.com/kudryashov216/goWebSock/internal"

	"github.com/kudryashov216/goWebSock/internal/websock"
	"github.com/kudryashov216/goWebSock/internal/websockclient"
)

type (
	RequestRW      internal.RequestRW
	WebSock        websock.WebSock
	WebSockClient  websockclient.WebSockClient
	ContextWebSock websock.ContextWebSock
)
