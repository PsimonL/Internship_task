/*
Tests for openx.go
*/

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertFahrenheitToCelsiusSuccess(t *testing.T) {
	/*
    Test conversion of Fahrenheit to Celsius with successful response.
	*/
	reqBody := map[string]interface{}{"fahrenheit": 32}
	reqBodyBytes, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(convertFahrenheitToCelsius)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

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
	/*
    Test conversion of 32°F to 0°C.
	*/
	reqBody := map[string]interface{}{"fahrenheit": 32}
	reqBodyBytes, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(convertFahrenheitToCelsius)
	handler.ServeHTTP(rr, req)

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
	/*
	Test conversion of 212°F to 100°C.
    */
	reqBody := map[string]interface{}{"fahrenheit": 212}
	reqBodyBytes, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(convertFahrenheitToCelsius)
	handler.ServeHTTP(rr, req)

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
