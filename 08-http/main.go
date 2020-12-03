package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// our request struct
type greetRequest struct {
	Name string
}

// our response struct
type greetResponse struct {
	Message string
}

func main() {
	// first, we need to create our route handler
	handler := http.NewServeMux()

	// then tell it which routes to handle, and how.
	// in this case we're instructing it to handle a route matching
	// `/greet`. A handler function receives a http.ResponseWriter and a pointer
	// to a http.Request. If a match occurs, the function is called.
	handler.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		// here we're doing a basic check on the http method of the request
		// we're only accepting POST requests for this handler
		if r.Method != http.MethodPost {
			// if it's not a POST request, we return MethodNotAllowed
			w.WriteHeader(http.StatusMethodNotAllowed)
			response := &greetResponse{
				Message: "method not allowed",
			}
			// and encode our response to JSON while writing to the ResponseWriter
			json.NewEncoder(w).Encode(&response)
			return
		}

		var payload greetRequest
		// here we're decoding the request body from JSON to our request struct
		json.NewDecoder(r.Body).Decode(&payload)

		// setting the content-type response header
		w.Header().Set("Content-Type", "application/json")
		// setting our response status code
		w.WriteHeader(http.StatusOK)
		response := &greetResponse{
			Message: fmt.Sprintf("hello %s! Nice to meet you!", payload.Name),
		}
		// encoding our response to JSON and writing to the responseWriter
		json.NewEncoder(w).Encode(response)
	})

	// here we create out http server, setting it's configurations and passing
	// our handler to it.
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	// a simple log message to indicate that our server is running
	log.Printf("http server starting on port %s", srv.Addr)

	// and here we startup our server.
	// this will lock the process until it's stopped.
	// to test this program out, after you run it, open a different terminal window
	// and use: curl --header "Content-Type: application/json" --request POST --data '{"name":"John Doe"}' http://localhost:8080/greet
	// this will issue a POST request to the /greet handler, passing in a JSON request body
	log.Fatal(srv.ListenAndServe())
}
