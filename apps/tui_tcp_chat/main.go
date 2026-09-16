package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const listenAddr = ":9000"

type client struct {
	conn net.Conn
	send chan string
}

type message struct {
	sender *client
	text   string
}

func main() {
	register := make(chan *client)
	unregister := make(chan *client)
	broadcast := make(chan message)

	go runBroadcaster(register, unregister, broadcast)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("listen %s: %v", listenAddr, err)
	}
	log.Printf("tui_tcp_chat server listening on %s", listenAddr)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		ln.Close()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Printf("listener closed: %v", err)
				return
			}
			log.Printf("accept error: %v", err)
			continue
		}
		go handleConn(conn, register, unregister, broadcast)
	}
}

// runBroadcaster is the single goroutine that owns the set of connected
// clients. All membership changes and message fan-out happen here, over
// channels — no mutex needed since only this goroutine ever touches the map.
func runBroadcaster(register, unregister chan *client, broadcast chan message) {
	clients := make(map[*client]bool)
	for {
		select {
		case c := <-register:
			clients[c] = true

		case c := <-unregister:
			if _, ok := clients[c]; ok {
				delete(clients, c)
				close(c.send)
			}

		case m := <-broadcast:
			for c := range clients {
				if c == m.sender {
					continue
				}
				select {
				case c.send <- m.text:
				default:
					// c's send buffer is full: drop the message for this
					// client rather than block delivery to everyone else.
				}
			}
		}
	}
}

func handleConn(conn net.Conn, register, unregister chan *client, broadcast chan message) {
	remote := conn.RemoteAddr()
	log.Printf("client connected: %s", remote)

	c := &client{conn: conn, send: make(chan string, 16)}
	register <- c

	writerDone := make(chan struct{})
	go func() {
		defer close(writerDone)
		for msg := range c.send {
			if _, err := fmt.Fprintln(conn, msg); err != nil {
				conn.Close()
				return
			}
		}
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		broadcast <- message{sender: c, text: text}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("read error from %s: %v", remote, err)
	}

	unregister <- c
	<-writerDone
	conn.Close()
	log.Printf("client disconnected: %s", remote)
}
