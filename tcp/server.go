package main

import (
	"log"
	"net"
)

func main() {
	network := "tcp"
	address := ":8080"
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		log.Printf("resolve %s://%s fail %v", network, address, err)
	}
	conn, err := net.ListenTCP(network, addr)
	if err != nil {
		log.Printf("create connection failed %v", err)
	}
	log.Printf("start listen")
	for {
		listener, err := conn.AcceptTCP()
		if err != nil {
			log.Printf("listen failed %v", err)
		}
		buffer := make([]byte, 64)
		n, err := listener.Read(buffer)
		if err != nil {
			log.Printf("read fail")
		}
		ping := string(buffer[:n])
		log.Printf("message content %s %d", ping, n)
		if ping == "ping" {
			m, err := listener.Write([]byte("pong"))
			if err != nil {
				log.Printf("write fail")
			}
			log.Printf("write %d bytes", m)
		} else {
			log.Printf("ping content not match")
		}
	}
}
