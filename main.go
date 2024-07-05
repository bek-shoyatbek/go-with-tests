package main

import (
	"log"
	"net/http"
	httpserver "playground/http-server"
)

func main() {
	server := &httpserver.PlayerServer{Store: httpserver.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
