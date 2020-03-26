package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	network := "udp"
	address := ":8080"

	addr := resolveAddress(network, address)

	conn, err := net.ListenUDP(network, addr)
	if err != nil {
		log.Printf("establish connection fail %v", err)
	}
	for {
		buffer := make([]byte, 512)
		n, clientAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("read error %v", err)
		}
		log.Printf("read %d bytes from %v", n, clientAddr)
		msg := string(buffer[:n])

		m, err := conn.WriteToUDP([]byte(strings.ToUpper(msg)), resolveAddress(clientAddr.Network(), clientAddr.String()))
		if err != nil {
			log.Printf("write error %v", err)
		}
		log.Printf("write %d bytes to %v", m, clientAddr)
	}

}

func resolveAddress(network string, address string) *net.UDPAddr {
	addr, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		log.Printf("resovle address error %v", err)
	}
	return addr
}
