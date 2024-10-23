package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"signing-service-challenge/api"
	"signing-service-challenge/persistence"
)

// TestCreateDevice tests the device creation handler
func TestCreateDevice(t *testing.T) {
	store := persistence.NewInMemorys()
	deviceHandler := api.NewDeviceHandler(store)

	// Define test cases
	tests := []struct {
		name       string
		apiRoute   string
		handler    http.HandlerFunc
		body       string // Add body here for POST requests
		expectCode int
	}{
		{
			name:       "Create Device",
			apiRoute:   "/api/v0/devices",
			handler:    deviceHandler.CreateDevice,
			body:       `{"id": "device1", "algorithm": "RSA", "label": "Test Device"}`, // Example body
			expectCode: http.StatusCreated,
		},
		{
			name:       "List Devices",
			apiRoute:   "/api/v0/devices/list",
			handler:    deviceHandler.CreateDevice,                                      // Modify if needed
			body:       `{"id": "device1", "algorithm": "RSA", "label": "Test Device"}`, // Example body
			expectCode: http.StatusCreated,
		},
	}

	// Iterate through test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a request with the provided body
			req, err := http.NewRequest(http.MethodPost, test.apiRoute, bytes.NewBufferString(test.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json") // Set Content-Type header

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler with the ResponseRecorder and request
			test.handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != test.expectCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectCode)
			}
		})
	}
}
