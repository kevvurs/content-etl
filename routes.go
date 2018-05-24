package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// Configures and returns a server instance
func buildRouter() *mux.Router {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	return mx
}

// Initializes the app's HTTP routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/etl", etlHandler(formatter)).Methods("GET")
	mx.HandleFunc("/clean", destroyHandler(formatter)).Methods("GET")
	mx.HandleFunc("/_ah/warmup", pingHandler(formatter)).Methods("GET")
}
