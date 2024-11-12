package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandler checks different scenarios for the handler function.
func TestHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		body       string
		statusCode int
	}{
		{
			name:       "GET request to root path",
			method:     http.MethodGet,
			path:       "/",
			statusCode: http.StatusOK,
		},
		{
			name:       "GET request to ascii-art",
			method:     http.MethodGet,
			path:       "/ascii-art",
			statusCode: http.StatusOK,
		},
		{
			name:       "POST request with valid input",
			method:     http.MethodPost,
			path:       "/ascii-art",
			body:       "userText=Hello&style=standard",
			statusCode: http.StatusOK,
		},
		{
			name:       "POST request with invalid input",
			method:     http.MethodPost,
			path:       "/",
			body:       "userText=Invalid€€Input&style=standard",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "POST request with unavailable banner",
			method:     http.MethodPost,
			path:       "/ascii-art",
			body:       "userText=Hello&style=ariel",
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Invalid path",
			method:     http.MethodGet,
			path:       "/invalid",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "Invalid method",
			method:     http.MethodHead,
			path:       "/",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creating a request
			req := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			// Telling the server how the data is encoded
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Creating a response recorder to record the response
			ResponseRecorder := httptest.NewRecorder()

			// Calling the handler with the ResponseRecorder and request
			handler(ResponseRecorder, req)

			// Checking if the status code matches the expected status code
			if ResponseRecorder.Code != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					ResponseRecorder.Code, tt.statusCode)
			}
		})
	}
}
