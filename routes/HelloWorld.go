package routes

import (
	"io"
	"log"
	"net/http"
)

/*
Basic hello world function that recieves a request from the http server in net/http and writes out
that a request was recieved and logged to the console and sends back "Hello World!" to the requester.
*/
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved Request at /HelloWorld")

	io.WriteString(w, "Hello World!")
}
