package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Book struct {
	ID     string
	Name   string
	Author string
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetails(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	raw, readerr := ioutil.ReadFile("books.json")

	if readerr != nil {
		log.Println("error", readerr)
		os.Exit(1)
	}
	//Unmarshal JSON raw data into books
	marshallerr := jsonparse.Unmarshal(raw, &books)
	if marshallerr != nil {
		log.Fatal(marshallerr)
	}
	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	//Create rpc server
	s := rpc.NewServer()
	//Register type of data
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}

//Client- Can be curl
//Start server and run curl

//curl -X POST \
// http://localhost:1234/rpc \
// -H 'cache-control: no-cache' \
// -H 'content-type: application/json' \
// -d '{
// "method": "JSONServer.GiveBookDetail",
// "params": [{ "ID": "1234" }], "id": "1" }'
