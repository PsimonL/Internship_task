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
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func convertFahrenheitToCelsius(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request body"})
		return
	}

	fahrenheit, ok := requestBody["fahrenheit"].(float64)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid Fahrenheit temperature"})
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
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/convert", convertFahrenheitToCelsius)
	fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
