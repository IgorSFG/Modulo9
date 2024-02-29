package main

import (
	"testing"
	"io/ioutil"
	"net/http"
)

func testGet(t *testing.T) {
	url := "http://localhost:8000/sensors"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		t.Errorf("Error making GET request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	t.Logf("Response body: %s", body)
}

func TestRequests(t *testing.T) {
	testGet(t)
}
