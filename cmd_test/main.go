package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/kudryashov216/goWebSock"
)

func main() {

	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	defer listener.Close()

	for {

		conn, err := listener.Accept()

		if err == io.EOF {
			fmt.Println("Клиент разорвал соединение!")
			continue
		}

		goWebSock.CreateWebSockServer(conn)

	}

}
