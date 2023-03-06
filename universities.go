package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Json formatting for requested data from API
type University struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	AlphaTwo string   `json:"alpha_two_code"`
	WebPages []string `json:"web_pages"`
	Map      struct {
		OpenStreetMaps string `json:"openstreetmaps"`
	} `json:"maps"`
	Languages map[string]string `json:"languages"`
}

// uniRequest is a function that sends HTTP GET request to the universities API
func uniRequest(endpoint string, params map[string]string) (*http.Response, error) {

	// Construct the API URL with the endpoint and query parameters.
	apiUrl := "http://universities.hipolabs.com/" + endpoint
	queryString := url.Values{}
	for k, v := range params {
		queryString.Set(k, v)
	}

	apiUrl += "?" + queryString.Encode()
	// Send the GET request and handle any errors.
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// searchSimilarUniversities is a function that searches for universities with similar names in neighboring countries of a given country.
func searchSimilarUniversities(name string, countryCode string) ([]University, error) {
	var results []University

	// Get all universities in neighboring countries
	country, err := getCountry(countryCode)
	var borderCodes = country.Borders

	if err != nil {
		return results, err
	}

	// Search for universities with similar names in neighboring countries
	for _, country := range borderCodes {
		// Build the URL to search for universities with the given name in the current neighboring country
		url := fmt.Sprintf("http://universities.hipolabs.com/search?country=%s&name=%s", country, name)
		// Send a request to the universities API to get the universities with the given name in the current neighboring country
		resp, err := http.Get(url)
		if err != nil {
			return results, err
		}
		// Parse the response to get the list of universities
		var universities []University
		if err := json.NewDecoder(resp.Body).Decode(&universities); err != nil {
			return results, err
		}
		// Add universities with similar names to the results
		for _, university := range universities {
			if university.Name != name {
				results = append(results, university)
			}
		}
	}
	// Return the list of universities with similar names in neighboring countries
	return results, nil
}

// Handler function for /universities/{name} endpoint
func uniEndpoint(w http.ResponseWriter, r *http.Request) {

	// Parse the language from the request path
	name := r.URL.Query().Get("name")
	country := r.URL.Query().Get("country")
	searchUniByNameAndCountry(name, country)
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

		similarUniversities, err := searchSimilarUniversities(name, country)
		if err != nil {
			http.Error(w, "Error fetching university data", http.StatusInternalServerError)
			log.Printf("Error fetching university data: %v", err)
			return

			uniResults := append(universities, similarUniversities...)

			json.NewEncoder(w).Encode(uniResults)

		}
	}

	// Return the university details in JSON format
	json.NewEncoder(w).Encode(universities)
}
