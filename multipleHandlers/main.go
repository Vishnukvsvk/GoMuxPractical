package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		a := rand.Intn(100)
		io.WriteString(w, strconv.Itoa(a))
	})

	mux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	http.ListenAndServe(":8080", mux)
}
