package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handles the requests to the root URL
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Retreive request information
	method := r.Method                      // Get the HTTP method (GET, POST, etc.)
	url := r.URL.String()                   // Get the requested URL
	headers := r.Header                     // Get the request headers
	userAgent := r.Header.Get("User-Agent") // Get the User-Agent header (info about the client - browser or HTTP client)

	// Create a message using the request information
	message := fmt.Sprintf(
		"Welcome to the Real-Time Messenger!\n"+
			"Request Method: %s\n"+
			"Request URL: %s\n"+
			"User-Agent^ %s\n", method, url, userAgent)

	// Add all headers to the message
	message += "Headers:\n"
	for key, values := range headers {
		for _, value := range values {
			message += fmt.Sprintf("%s: %s\n", key, value)
		}
	}

	// Write the message back to the client
	fmt.Fprint(w, message)
}
