package main

import (
	"log"

	"github.com/tmp-ticket/backend/routes"
)

func main() {

	server := routes.ServerSetup()

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}

}
