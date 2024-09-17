package routes

import (
	"io"
	"log"
	"net/http"
)

/*
Basic hello world function that recieves a request from the http server in net/http and writes out
that a request was recieved and logged to the console and sends back "Hello World!" to the requester
*/
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved Request on /HelloWorld")

	io.WriteString(w, "Hello World!")
}

/*
Route struct that contains a path string and handler function for http requests.
Is to be used inside of ServerSetup for convenience.
*/
type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

// Convenient list of routes to funuctions.
// DO NOT ADD ROUTES WHILE SERVER IS RUNNING!
var routes = []Route{
	{"/HelloWorld", HelloWorld},
}

// Convenience function to create an http server with routes needed for API.
// The Addr parameter is supposed to be a string that represents ip:port
// such as :8080 for port 8080 at localhost or 127.0.0.1:8080 for the same effect.
func ServerSetup(Addr string) *http.Server {

	mux := http.NewServeMux()

	// Adds all Route structs to the above mux for the http server.
	for _, route := range routes {
		mux.HandleFunc(route.Path, route.Handler)
	}
	//declaration of basic settings for http server
	server := &http.Server{
		Addr:    Addr,
		Handler: mux,
	}
	return server
}
