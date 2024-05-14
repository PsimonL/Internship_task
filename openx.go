/*
Sample API as part of OpenX Intern Assesment.
*/

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func generateAppIdentifier(length int) string {
	/*
	Generates app identifier.
	*/
	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func convertFahrenheitToCelsius(w http.ResponseWriter, r *http.Request) {
	/*
    Converts Fahrenheit temperature to Celsius.
    Receives JSON data with Fahrenheit temperature in the request body.
    Returns JSON response with the converted Celsius temperature.
	*/
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request body"}); err != nil { 
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	fahrenheit, ok := requestBody["fahrenheit"].(float64)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid Fahrenheit temperature"}); err != nil { 
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	celsius := (fahrenheit - 32) * 5.0 / 9.0
	appIdentifier := generateAppIdentifier(10)

	response := map[string]interface{}{
		"celsius":        celsius,
		"app_identifier": appIdentifier,
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Request will be served.")
	if err := json.NewEncoder(w).Encode(response); err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func probeHandler(w http.ResponseWriter, r *http.Request) {
	/*
	Satisfies K8S health, probe system.
	*/
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "K8s Probe request",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/convert", convertFahrenheitToCelsius)
	http.HandleFunc("/probe", probeHandler)
	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
		return
	}
}
