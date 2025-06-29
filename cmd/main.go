package main

import (
	"fmt"
	"log"
	"net/http"
)

type ApiHandler struct{}

func (ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := "Moses from the api"
	w.Write([]byte(name))
}

func main() {
	fmt.Println("this siw working")
	// server

	mux := http.NewServeMux()

	mux.Handle("/api", ApiHandler{})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	fmt.Println("Buda server iko on", 8000)
	log.Fatal(server.ListenAndServe())

}
