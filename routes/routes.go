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

func ServerSetup() *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorld)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return server
}
