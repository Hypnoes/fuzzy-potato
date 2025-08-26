package main

import (
	"log"
	"net/http"

	"fuzzypotato/api_manager"
	"fuzzypotato/api_service"
)

func main() {
	registry := &api_manager.ServiceRegistry{Service: make(map[string]api_service.ServiceConfig)}
	http.HandleFunc("/register", registry.Register)
	http.HandleFunc("/lookup", registry.Lookup)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
