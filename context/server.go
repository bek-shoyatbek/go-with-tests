package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			log.Fatal("cannot fetch store data")
		}
		fmt.Fprintf(w, data)
	}
}
