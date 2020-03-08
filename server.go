package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)



var upgrader = websocket.Upgrader{} // use default options
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	var addr = flag.String("addr", ":9000", "http service address")
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
		fmt.Printf("user default port %s", port)
	}

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
