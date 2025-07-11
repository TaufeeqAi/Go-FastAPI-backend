package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8000/ai/hello")
	if err != nil {
		http.Error(w, "Failed to contact AI service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from AI service", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Response from AI service: %s", string(body))
}
