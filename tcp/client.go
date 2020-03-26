package main

import (
	"log"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	start := time.Now().UnixNano()
	n, err := conn.Write([]byte("ping"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("write %d bytes", n)
	buffer := make([]byte, 4)
	m, err := conn.Read(buffer)
	if m != n {
		log.Println("content error")
	}
	if err != nil {
		log.Println(err)
	}
	if string(buffer) == "pong" {
		log.Printf("ping took %dms\n", (time.Now().UnixNano()-start)/int64(time.Millisecond))
	}
}
