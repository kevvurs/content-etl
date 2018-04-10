package service

import (
	"github.com/GoIncremental/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
)

// Configures and returns a server instance
func NewServer() *negroni.Negroni {
	log.Println("Launching Go server")
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// Initializes the app's HTTP routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
}
