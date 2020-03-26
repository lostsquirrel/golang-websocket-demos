package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

func main() {
	origin := "http://d6bc7dd3a64245c0baae023d61cf84d6-cn-hangzhou.alicloudapi.com"
	url := "ws://d6bc7dd3a64245c0baae023d61cf84d6-cn-hangzhou.alicloudapi.com/register"
	url = "ws://localhost:9000"
	ws, err := websocket.Dial(url, "chat", origin)
	if err != nil {
		log.Print("connect error")
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("RG#ffd3234343dae324342@12344133")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}