package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Json formatting for requested data from API
type Country struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName struct {
			Nno struct {
				Official string `json:"official"`
				Common   string `json:"common"`
			} `json:"nno"`
		} `json:"nativeName"`
	} `json:"name"`
	Tld         []string `json:"tld"`
	Cca2        string   `json:"cca2"`
	Ccn3        string   `json:"ccn3"`
	Cca3        string   `json:"cca3"`
	CIOC        string   `json:"cioc"`
	Independent bool     `json:"independent"`
	Status      string   `json:"status"`
	UnMember    bool     `json:"unMember"`
	Currencies  struct {
		NOK struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"NOK"`
	} `json:"currencies"`
	Idd struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
	Capital   []string `json:"capital"`
	Region    string   `json:"region"`
	Subregion string   `json:"subregion"`
	Languages struct {
		Nno string `json:"nno"`
		Nob string `json:"nob"`
		Smi string `json:"smi"`
	} `json:"languages"`
	Latlng     []float64 `json:"latlng"`
	Landlocked bool      `json:"landlocked"`
	Borders    []string  `json:"borders"`
	Area       float64   `json:"area"`
	Demonyms   struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
	} `json:"demonyms"`
	Flag string `json:"flag"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	Population int `json:"population"`
	Gini       struct {
		Year2018 float64 `json:"2018"`
	} `json:"gini"`
	Timezones   []string `json:"timezones"`
	Continents  []string `json:"continents"`
	StartOfWeek string   `json:"startOfWeek"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode"`
}

// countryRequest sends an HTTP GET request to the restcountries.com API with the provided endpoint and country code.
func countryRequest(endpoint string, countryCode string) (*http.Response, error) {
	apiUrl := fmt.Sprintf("https://restcountries.com/v3/%s/%s", endpoint, countryCode)
	resp, err := http.Get(apiUrl)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Handler function for /countries/{countryCode} endpoint
func getCountry(countryCode string) (Country, error) {
	var results Country
	// Send a request to the countries API to get the country's name
	resp, err := countryRequest("alpha", countryCode)

	if err != nil {
		return results, err
	}

	defer resp.Body.Close()

	var countries []Country
	err = json.NewDecoder(resp.Body).Decode(&countries)
	if err != nil {
		panic(err)
	}

	results = countries[0]

	return results, err
}

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
