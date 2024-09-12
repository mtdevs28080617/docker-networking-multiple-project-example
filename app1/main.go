package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response structure for JSON output
type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := Response{Message: "ping"}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	log.Fatalln(http.ListenAndServe(":80", nil))
}
