package main

// Div tester og restkode
/*
func TestUniInfoHandler(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/universities/Harvard%20University", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(uniInfoHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `{"name":"Harvard University","country":"United States","alpha_two_code":"US","web_pages":["http://www.harvard.edu/"],"languages":{"english":"English"},"map":"https://www.openstreetmap.org/relation/165640"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCountryUnisHandler(t *testing.T) {
	// Create a request with the desired country code
	req, err := http.NewRequest("GET", "/countries/US/universities", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the response recorder and request
	handler := http.HandlerFunc(countryUnisHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	expectedContentType := "application/json"
	if ct := rr.Header().Get("Content-Type"); ct != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v", ct, expectedContentType)
	}

	// Check the body of the response
	expectedBody := `{"country":"United States","universities":["University of California, Los Angeles (UCLA)","Harvard University","Stanford University","Massachusetts Institute of Technology (MIT)","California Institute of Technology (Caltech)","University of Chicago","University of Pennsylvania","Princeton University","Columbia University","Yale University","Duke University","Johns Hopkins University","New York University (NYU)","University of Michigan-Ann Arbor","Cornell University","Northwestern University","Brown University","University of Wisconsin-Madison","University of California, San Diego (UCSD)","University of California, Berkeley (UCB)","University of Illinois at Urbana-Champaign (UIUC)","University of California, Santa Barbara (UCSB)","University of Minnesota Twin Cities"]}`
	if body := strings.TrimSuffix(rr.Body.String(), "\n"); body != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedBody)
	}

}
*/
