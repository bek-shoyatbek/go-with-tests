package main

import (
	"log"
	"net/http"
	httpserver "playground/http-server"
)

func main() {
	handler := http.HandlerFunc(httpserver.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
