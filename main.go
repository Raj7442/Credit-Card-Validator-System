package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func luhnAlgorithm(cardNumber string) bool {
	var sum int
	alt := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}
	return sum%10 == 0
}

func main() {
	http.HandleFunc("/validate", validateCardHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

type RequestPayload struct {
	CardNumber string `json:"card_number"`
}

type ResponsePayload struct {
	Valid bool `json:"valid"`
}

func validateCardHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON request
	var payload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || len(payload.CardNumber) == 0 {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validate the credit card number
	isValid := luhnAlgorithm(payload.CardNumber)

	// Create JSON response
	response := ResponsePayload{Valid: isValid}

	// Send the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
