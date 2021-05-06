package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func serveTCP2(conn *websocket.Conn) {
	l, err := net.Listen("tcp4", ":9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())
	for {
		c, err := l.Accept()
		log.Println("get a new client")
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println(conn)
		c.Write([]byte("connected"))
		conn.Write([]byte("connected"))
		go func() {
			io.Copy(conn, c)
		}()
		io.Copy(c, conn)
	}

}

func main() {

	http.Handle("/echo", websocket.Handler(func(conn *websocket.Conn) {
		log.Println("start a new worker")
		 serveTCP2(conn)
	}))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
