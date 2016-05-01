package main

import (
    "net/http"
    "log"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", basicHandler)

    log.Panic(http.ListenAndServe("localhost:1337", router))
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world!"))
}
