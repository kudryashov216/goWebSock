package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kudryashov216/goWebSock"
)

func main() {

	// listener, err := net.Listen("tcp", ":8000")

	// if err != nil {
	// 	log.Fatalf("%s\n", err.Error())
	// }

	// defer listener.Close()

	http.HandleFunc("/chat", Handler)
	http.ListenAndServe(":8001", nil)

	// for {

	// 	conn, err := listener.Accept()

	// 	go func(c net.Conn) {

	// 		ws, err := goWebSock.CreateWebSockServer(c, "")

	// 		if err != nil {
	// 			log.Println("handshake error:", err)
	// 			return
	// 		}

	// 		defer ws.Close()

	// 		buf := make([]byte, 4096)

	// 		for {
	// 			n, err := c.Read(buf)
	// 			if err != nil {
	// 				log.Println("connection closed:", err)
	// 				return
	// 			}
	// 			log.Printf("got %d bytes: %b", n, buf[:n])
	// 		}

	// 	}(conn)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}

	// }

}

func Handler(rw http.ResponseWriter, r *http.Request) {

	var buff = make([]byte, 4096)

	var reqRW goWebSock.RequestRW

	reqRW.Request = r
	reqRW.ResponseWriter = rw

	webSock, err := goWebSock.CreateWebSockServerFromRequestRW(&reqRW)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = webSock.Send("Hi. This is websocket is written in the Go programming language. Welcome!")

	if err != nil {
		log.Println(err.Error())
		webSock.Close()
	}

	for {

		webSock.Read(&buff)

		if len(buff) > 0 {
			webSock.Send("Pong!")
		}

	}

}
