package main

import (
	"encoding/json"
	"fmt"
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

/*
To do list

1.  Lage en getUniversity basert på navn og/eller land

2. Search similar universities.. Måhuske å legge inn navn og landekode for første uni. dette må legges inn i api søket i postman

*/
