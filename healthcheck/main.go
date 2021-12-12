package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/health", healthCheck)
	port := ":8080"
	log.Fatal(http.ListenAndServe(port, nil))
}
