package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conn map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conn: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWSDataStream(ws *websocket.Conn) {
	fmt.Println("New connection to data stream: ", ws.RemoteAddr())

	// since this is only attached and the message is not read hence the data stream is sent to all clients
	for {
		payload := fmt.Sprint("Current time", time.Now().String())
		ws.Write([]byte(payload))
		time.Sleep(time.Second)
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New connection: ", ws.RemoteAddr())

	// make a mutex to protect parallel access to conn
	s.conn[ws] = true

	// read the messages in a seperate goroutine loop
	s.readloop(ws)
}

func (s *Server) readloop(ws *websocket.Conn) {
	buf := make([]byte, 4096)
	for {
		n, err := ws.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed: ", ws.RemoteAddr())
				// remove the connection from the map
				delete(s.conn, ws)
				break
			}

			fmt.Println("Error reading from client: ", err)
			// make a choice whether to continue or break the read loop
			continue
		}

		msg := buf[:n] // trim the buffer to the number of bytes read
		fmt.Println("Received Message", string(msg))
		// send message to only one client
		// ws.Write([]byte("Thanks for the message"))

		// send message to all clients
		// s.broadcast([]byte(fmt.Sprint("Thanks for the message", string(msg))))

		// send to all clients except the sender
		s.broadcastExceptSender([]byte(fmt.Sprint("Thanks for the message", string(msg))), ws)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conn {

		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Error writing to client: ", err)
			}
		}(ws)

	}
}

func (s *Server) broadcastExceptSender(b []byte, sender *websocket.Conn) {
	for ws := range s.conn {
		// ignore sending message to the sender
		if ws == sender {
			continue
		}

		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Error writing to client: ", err)
			}
		}(ws)

	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/stream", websocket.Handler(server.handleWSDataStream))

	fmt.Println("Listening on port ws://localhost:3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
