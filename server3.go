package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)


func main() {
	http.Handle("/echo", websocket.Handler(func(conn *websocket.Conn) {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Print(err)
		}
		content := string(buffer[:n])
		fmt.Printf("read %d bytes (%s)\n", n, content)
		n, err = conn.Write([]byte(content))
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("write %d bytes\n", n)
	}))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}