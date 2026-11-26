package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/kudryashov216/goWebSock"
	"github.com/kudryashov216/goWebSock/opcodes"
)

func main() {

	// listener, err := net.Listen("tcp", ":8000")

	// if err != nil {
	// 	log.Fatalf("%s\n", err.Error())
	// }

	// defer listener.Close()

	http.HandleFunc("/chat", Handler)
	http.ListenAndServe(":8000", nil)

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

	reqRW := goWebSock.NewRequestRW(rw, r)

	ctx, webSock, err := goWebSock.CreateWebSockServerFromRequestRW(reqRW)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go Reader(webSock)
	go Sender(ctx.GetContext(), webSock)

}

func Reader(ws goWebSock.WebSock) {

	for {

		frame, err := ws.ReadFrames()

		if err != nil {
			break
		}

		data, _ := frame.Unmask(frame.ReadFrame())

		fmt.Println(string(*data))

	}

}

func Sender(ctx context.Context, ws goWebSock.WebSock) {

	fileArray := [2]string{"C:\\Users\\Nkudryashov\\Desktop\\web socket.png", "C:\\Users\\Nkudryashov\\Desktop\\TaskManager.png"}

	for _, file := range fileArray {

		bytes, _ := os.ReadFile(file)

		_, err := ws.Send(opcodes.BINARY_FILE, bytes)

		if err != nil {
			ws.Close()
			return
		}

		_, err = ws.SendToString("Hello")

		if err != nil {
			ws.Close()
			return
		}

	}

	// reader := bufio.NewReader(os.Stdin)

	// type readResult struct {
	// 	line string
	// 	err  error
	// }
	// lineCh := make(chan readResult, 1)

	// for {
	// 	// читаем в отдельной горутине — чтение блокирующее, его не отменить
	// 	go func() {
	// 		line, err := reader.ReadString('\n')
	// 		lineCh <- readResult{line, err}
	// 	}()

	// 	select {
	// 	case <-ctx.Done():
	// 		// контекст отменили извне — корректно закрываем соединение и выходим
	// 		ws.SoftClose("bye")
	// 		return

	// 	case r := <-lineCh:
	// 		if r.err != nil {
	// 			if errors.Is(r.err, io.EOF) {
	// 				fmt.Println("stdin закрыт — закрываем соединение")
	// 				ws.SoftClose("bye")
	// 				return
	// 			}
	// 			fmt.Println("ошибка чтения stdin:", r.err)
	// 			return
	// 		}

	// 		line := strings.TrimSpace(r.line)
	// 		if line == "" {
	// 			continue
	// 		}
	// 		if line == "/quit" {
	// 			ws.SoftClose("bye")
	// 			return
	// 		}

	// 		if _, err := ws.SendToString(line); err != nil {
	// 			fmt.Println("ошибка отправки:", err)
	// 			return
	// 		}
	// 	}
	// }

}
