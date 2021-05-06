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

type Wrapper struct {
	C websocket.Conn
}

func serveTCP(conns chan *websocket.Conn) {
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
		go func() {
			log.Printf("work number %d\n", len(conns))
			conn := <-conns
			log.Println(conn)
			log.Println(conn.Config())
			c.Write([]byte("connected"))
			conn.Write([]byte("connected"))
			go func() {
				io.Copy(conn, c)
			}()
			io.Copy(c, conn)
		}()

	}
}

func serveTick(conns chan Wrapper) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	log.Printf("connections %d", len(conns))
	cw := <-conns
	conn := &cw.C
	buffer := make([]byte, 1024)
	for {
		msg := fmt.Sprintf("%v", <-ticker.C)
		log.Printf("sned %s\n", msg)
		//log.Println(conn != nil)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("game over")
			log.Fatal(err)
		}
		//log.Println(conn.MaxPayloadBytes)
		//log.Println(conn.PayloadType)
		//log.Println(conn.IsClientConn())
		//log.Println(conn.IsServerConn())
		////log.Println(conn.Len())
		//log.Println(conn.LocalAddr())
		//log.Println(conn.RemoteAddr())
		//log.Println(conn.Request())
		//log.Println(conn.Config())
		//log.Println(conn.HeaderReader())
		read(conn, buffer)
	}

}

func read(conn *websocket.Conn, buffer []byte) {
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Print(err)
	}
	content := string(buffer[:n])
	fmt.Printf("read %d bytes (%s)\n", n, content)
}

func main() {
	conns := make(chan Wrapper, 6)
	buffer := make([]byte, 1024)
	//c2 := make(map[string]*websocket.Conn)

	go serveTick(conns)
	http.Handle("/echo", websocket.Handler(func(conn *websocket.Conn) {
		log.Println("start a new worker")
		conn.Write([]byte("accepted"))
		err := conn.SetDeadline(time.Now().Add(1 * time.Minute))
		if err != nil {
			fmt.Print(err)
		}
		//conn.PayloadType = websocket.ContinuationFrame
		read(conn, buffer)
		conn.Write([]byte("go go go \n"))
		conns <- Wrapper{*conn}

	}))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
