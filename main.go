package main

import (
	"log"

	"github.com/tmp-ticket/backend/routes"
)

func main() {

	const addr = ":8080"

	server := routes.ServerSetup(addr)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
