package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eriktate/NaaSgul/config"
	"github.com/eriktate/NaaSgul/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var wsMap map[int]*websocket.Conn
var counter int

type tester struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func main() {
	counter = 0
	wsMap = make(map[int]*websocket.Conn)

	router := mux.NewRouter()
	handlers.InitNotificationHandler(NotificationProvider(), router.PathPrefix("/api/notification").Subrouter())
	router.HandleFunc("/ws", WsHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	host := config.GetServerHost()
	port := config.GetServerPort()

	log.Printf("Starting server on %s:%s...", host, port)
	log.Panic(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	t := &tester{}
	err := decoder.Decode(t)

	if err != nil {
		log.Println(err)
	}

	wsMap[t.ID].WriteJSON(t)
}
