package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/inventory/routes"
)

type ApiHandler struct{}

func (ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := "Moses from the api"
	w.Write([]byte(name))
}

func main() {
	fmt.Println("this siw working")
	// server

	mux := routes.Routers()

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	fmt.Println("Buda server iko on", 8000)
	log.Fatal(server.ListenAndServe())

}
