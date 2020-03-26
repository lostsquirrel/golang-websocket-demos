package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	network := "udp"
	address := "localhost:8080"
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Printf("create connection to %s://%s failed", network, address)
	}
	log.Print("ready to communication to udp server!\n")
	for {
		buffer := make([]byte, 512)
		reader := bufio.NewReader(os.Stdin)
		n, err := reader.Read(buffer)
		if err != nil {
			log.Printf("read fail from console %s\n", err)
		}
		log.Printf("read %d bytes from console\n", n)
		m, err := conn.Write(buffer[:n])
		if err != nil {
			log.Printf("write fail %s", err)
		}
		log.Printf("write %d bytes to udp\n", m)
		x, err := conn.Read(buffer)
		if err != nil {
			log.Printf("read fail from udp %s\n", err)
		}
		log.Printf("%s\n", string(buffer[:x]))
	}

}
