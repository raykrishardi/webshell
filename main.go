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

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
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
