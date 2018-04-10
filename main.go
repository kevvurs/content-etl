package main

import (
	service "github.com/kevvurs/newsroom-content-persistence/service"
	"google.golang.org/appengine"
	"log"
	"net/http"
)

func main() {
	log.Println("Booting up newsroom-content-persistence")
	if router := service.BuildRouter(); router != nil {
		http.Handle("/", router)
	}
	appengine.Main()
}
