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

const portNumber = ":8081"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// reader will continuously listen for new messages being sent to the websocket endpoint
func reader(conn *websocket.Conn) {
	go func() {
		for {
			// read in a message
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			// print out the message for debugging
			fmt.Println(string(msg))

			// Get commands to be executed from the front-end
			commands := strings.Split(string(msg), " ")
			if len(commands) < 1 {
				log.Println("Empty command...")
				return
			}

			// Execute the command in its own pty (pseudo-terminal)
			c := exec.Command("websh", commands[1:]...)
			f, err := pty.Start(c)
			if err != nil {
				panic(err)
			}

			// Send the output of the command from pty back to the front-end
			var buf bytes.Buffer
			io.Copy(&buf, f)
			err = conn.WriteMessage(1, buf.Bytes())
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()
}

func xtermEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected")
	reader(ws)
}

func main() {
	http.HandleFunc("/xterm", xtermEndpoint)
	fmt.Printf("Server listening on port %s\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
