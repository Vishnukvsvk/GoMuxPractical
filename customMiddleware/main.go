package main

import (
	"fmt"
	"net/http"
)

func middleware(originalhandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!!")
		//Pass control back to handler
		originalhandler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler...")
	w.Write([]byte("OKK"))
}

func main() {
	originalhandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalhandler))
	http.ListenAndServe(":8080", nil)
}
