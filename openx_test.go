package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertFahrenheitToCelsiusSuccess(t *testing.T) {
	// Prepare a request body with Fahrenheit value
	reqBody := map[string]interface{}{"fahrenheit": 32}
	reqBodyBytes, _ := json.Marshal(reqBody)

	// Create a request with the prepared body
	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and ResponseRecorder
	handler := http.HandlerFunc(convertFahrenheitToCelsius)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	// Check if the response body contains the expected fields
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding JSON response: %v", err)
	}

	if _, ok := response["celsius"]; !ok {
		t.Error("response does not contain 'celsius' field")
	}

	if _, ok := response["app_identifier"]; !ok {
		t.Error("response does not contain 'app_identifier' field")
	}
}

func TestConvert32FTo0C(t *testing.T) {
	// Prepare a request body with Fahrenheit value
	reqBody := map[string]interface{}{"fahrenheit": 32}
	reqBodyBytes, _ := json.Marshal(reqBody)

	// Create a request with the prepared body
	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and ResponseRecorder
	handler := http.HandlerFunc(convertFahrenheitToCelsius)
	handler.ServeHTTP(rr, req)

	// Check if the response body contains the expected Celsius value
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding JSON response: %v", err)
	}

	if celsius, ok := response["celsius"].(float64); ok {
		if celsius != 0 {
			t.Errorf("incorrect Celsius value for 32°F: got %v want 0", celsius)
		}
	} else {
		t.Error("response does not contain valid 'celsius' field")
	}
}

func TestConvert212FTo100C(t *testing.T) {
	// Prepare a request body with Fahrenheit value
	reqBody := map[string]interface{}{"fahrenheit": 212}
	reqBodyBytes, _ := json.Marshal(reqBody)

	// Create a request with the prepared body
	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and ResponseRecorder
	handler := http.HandlerFunc(convertFahrenheitToCelsius)
	handler.ServeHTTP(rr, req)

	// Check if the response body contains the expected Celsius value
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding JSON response: %v", err)
	}

	if celsius, ok := response["celsius"].(float64); ok {
		if celsius != 100 {
			t.Errorf("incorrect Celsius value for 212°F: got %v want 100", celsius)
		}
	} else {
		t.Error("response does not contain valid 'celsius' field")
	}
}
