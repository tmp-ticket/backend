package routes

import (
	"fmt"
	"io"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved Request")

	io.WriteString(w, "Hello World!")
}

func ServerSetup(Addr string) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorld)

	server := &http.Server{
		Addr:    Addr,
		Handler: mux,
	}
	return server
}
