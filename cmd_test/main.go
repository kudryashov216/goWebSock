package main

import (
	"fmt"
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

		go func(c net.Conn) {

			ws, err := goWebSock.CreateWebSockServer(c)

			if err != nil {
				log.Println("handshake error:", err)
				return
			}

			defer ws.Close()

			buf := make([]byte, 4096)
			for {
				n, err := c.Read(buf)
				if err != nil {
					log.Println("connection closed:", err)
					return
				}
				log.Printf("got %d bytes: %x", n, buf[:n])
			}

		}(conn)

		if err != nil {
			fmt.Println(err)
			continue
		}

	}

}
