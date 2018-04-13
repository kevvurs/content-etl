package main

import (
	"google.golang.org/appengine"
	"log"
	"net/http"
)

func main() {
	log.Println("Booting up newsroom-content-persistence")
	if router := buildRouter(); router != nil {
		http.Handle("/", router)
	}
	appengine.Main()
}
