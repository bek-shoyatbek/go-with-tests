package main

import (
	"log"
	"net/http"
	json_server "playground/json-server"
)

func main() {
	server := &json_server.PlayerServer{Store: json_server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
