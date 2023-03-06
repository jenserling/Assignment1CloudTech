package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// searches for universities with a given name and country using the Hipolabs API
func searchUniByNameAndCountry(name, country string) ([]University, error) {
	// build the URL with the query parameters
	baseURL := "http://universities.hipolabs.com/search"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}
	q := u.Query()
	q.Set("name", name)
	q.Set("country", country)
	u.RawQuery = q.Encode()

	// send the request and parse the response
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	var universities []University
	if err := json.NewDecoder(resp.Body).Decode(&universities); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return universities, nil
}
