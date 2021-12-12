package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

type UUID struct{}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)   // make byte array of length c
	_, err := rand.Read(b) //Random bits in b byte array
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}

func main() {
	mux := &UUID{}
	http.ListenAndServe(":8080", mux)
}
