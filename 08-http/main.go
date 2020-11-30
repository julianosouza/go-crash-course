package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type greetRequest struct {
	Name string
}

type greetResponse struct {
	Message string
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			response := &greetResponse{
				Message: "method not allowed",
			}
			json.NewEncoder(w).Encode(&response)
			return
		}

		var payload greetRequest
		json.NewDecoder(r.Body).Decode(&payload)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := &greetResponse{
			Message: fmt.Sprintf("hello %s! Nice to meet you!", payload.Name),
		}
		json.NewEncoder(w).Encode(response)
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Printf("http server starting on port %s", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
