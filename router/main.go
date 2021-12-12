package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goversion)
	router.GET("/api/v1/show-file/:name", getFileContent) //:name is path parameter. Not supported in default router
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

//If you observe the code, we used /usr/local/go/bin/go as the Go executable location because it is the Go compiler location in Mac OS X. While executing exec.Command, you should give the absolute path of the executable. So, if you are working on an Ubuntu machine or Windows, use the path to your installed Go executable. On Linux machines, you can easily find that out by using the $ which go command.
func goversion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("/usr/local/go/bin/go", "version")
	io.WriteString(w, response)
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/bin/cat", params.ByName("name")))
}
