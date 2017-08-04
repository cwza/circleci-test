package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	log.Fatalf("failed to searve http: %+v", http.ListenAndServe(":"+"50011", nil))
}
