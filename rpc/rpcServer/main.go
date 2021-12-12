package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	//Fill the pointer
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)
	rpc.Register(timeServer)
	rpc.HandleHTTP()

	//Listen to request on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error", e)
	}
	http.Serve(l, nil)
}
