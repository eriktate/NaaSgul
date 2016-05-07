package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//WsHandler test
func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	wsMap[counter] = conn
	counter++
	log.Println("Received web socket connection: ", conn.RemoteAddr().String())

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Println(string(p))
		if err = conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
		}
	}
}
