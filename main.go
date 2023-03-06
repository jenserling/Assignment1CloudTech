package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "strconv"
	"strings"
)

func countryEndpoint(w http.ResponseWriter, r *http.Request) {
	// Parse the country code from the request path
	parts := strings.Split(r.URL.Path, "/")
	countryCode := strings.ToUpper(parts[len(parts)-2])

	country, err := getCountry(countryCode)

	if err != nil {
		http.Error(w, "Error fetching country data", http.StatusInternalServerError)
		log.Printf("Error fetching country data: %v", err)
		return
	}

	json.NewEncoder(w).Encode(country)
}

// Handler function for /universities/{name} endpoint
func uniEndpoint(w http.ResponseWriter, r *http.Request) {

	// Parse the language from the request path
	name := r.URL.Query().Get("name")
	country := r.URL.Query().Get("country")
	searchSimilarUniversities(name, country)
}

// Handler function for /universities/{name} endpoint
func uniEndpoints(w http.ResponseWriter, r *http.Request) {
	// Parse the query string
	name := r.URL.Query().Get("name")
	country := r.URL.Query().Get("country")
	showSimilar := r.URL.Query().Get("similar")

	// Get the university details from the API
	universities, err := searchUniByNameAndCountry(name, country)
	if err != nil {
		http.Error(w, "Error fetching university data", http.StatusInternalServerError)
		log.Printf("Error fetching university data: %v", err)
		return
	}

	if showSimilar == "true" {

		/* similarUniversities, err := searchSimilarUniversities(name, country)
		if err != nil {
			http.Error(w, "Error fetching university data", http.StatusInternalServerError)
			log.Printf("Error fetching university data: %v", err)
			return


			uniResults := nil // merger similarUniversities og universites


			json.NewEncoder(w).Encode(uniResults) */

		//Merge universities + similarUnis. og returer. searchSimilarUniversities

	}

	// Return the university details in JSON format
	json.NewEncoder(w).Encode(universities)
}

func main() {
	// Create the server and listen for requests
	port := "8080"
	log.Printf("Server listening on port %v", port)

	http.HandleFunc("/universities/", uniEndpoint)
	http.HandleFunc("/countries/", countryEndpoint)
	/*XX Noe Rart Her XXsearchUniByNameAndCountry*/ http.HandleFunc("/uniandcuntry/", uniEndpoints)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

// Structs for parsing the responses from the APIs

// Helper function to send requests to the APIs

/*
	To do list

	1. Hente universitet basert på navn.

	2. Søke etter land basert på universitetes landskode. For å hente hvilke naboland som finnes.

	3. Søker etter universitet med lignendene navn som også ligger i naboland.


*/

// Send a request to the universities API to get universities that teach the language
//resp, err := uniRequest("search", map[string]string{"language": language})
/*	if err != nil {
		http.Error(w, "Error fetching universities", http.StatusInternalServerError)
		log.Printf("Error fetching universities: %v", err)
		return
	}
	defer resp.Body.Close()

	// Parse the response and extract the university names and web pages
	var universities []University
	err = json.NewDecoder(resp.Body).Decode(&universities)
	if err != nil {
		http.Error(w, "Error parsing university data", http.StatusInternalServerError)
		log.Printf("Error parsing university data: %v", err)
		return
	}
	var unis []string
	for _, uni := range universities {
		unis = append(unis, uni.Name)
	}

	// Send the response
	response := map[string]interface{}{
		"language":     strings.Title(language),
		"universities": unis,
	}
	json.NewEncoder(w).Encode(response)

*/
