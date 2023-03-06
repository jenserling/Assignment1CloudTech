package main

import (
	"log"
	"net/http"
)

func main() {
	// Create the server and listen for requests
	port := "8080"
	log.Printf("Server listening on port %v", port)

	http.HandleFunc("/universities/", uniEndpoint)
	http.HandleFunc("/countries/", countryEndpoint)
	http.HandleFunc("/uniandcuntry/", uniEndpoints)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
