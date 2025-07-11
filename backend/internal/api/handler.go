package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AskHandler(w http.ResponseWriter, r *http.Request) {

	prompt := "Tell me a fun fact about the Go programming language."

	// Create the request body
	requestBody, err := json.Marshal(map[string]string{
		"prompt": prompt,
	})
	if err != nil {
		http.Error(w, "Failed to create request body", http.StatusInternalServerError)
		return
	}

	// Send the POST request to the AI service
	resp, err := http.Post("http://localhost:8000/ai/ask", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to contact AI service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read and return the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from AI service", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "AI Response: %s", string(body))
}
