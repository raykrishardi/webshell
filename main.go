package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/creack/pty"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	go func() {
		for {
			// read in a message
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			// print out that message for clarity
			fmt.Println(string(p))
			commands := strings.Split(string(p), " ")
			if len(commands) < 1 {
				log.Println("Empty command...")
				return
			}

			c := exec.Command("websh", commands[1:]...)
			f, err := pty.Start(c)
			if err != nil {
				panic(err)
			}
			var buf bytes.Buffer
			io.Copy(&buf, f)
			// resp := fmt.Sprintf("\n\n%s", string(buf.Bytes()))
			err = conn.WriteMessage(1, buf.Bytes())
			if err != nil {
				log.Println(err)
				return
			}

			// if err := conn.WriteMessage(messageType, p); err != nil {
			// 	log.Println(err)
			// 	return
			// }

		}
	}()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	// err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if err != nil {
	// 				log.Println(err)
	// }
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	setupRoutes()
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
