package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"signing-service-challenge/api"
	"testing"
)

func TestServer_Health(t *testing.T) {
	server := &api.Server{}

	// Test the Health endpoint with a GET request
	t.Run("Health GET request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		resp := httptest.NewRecorder()

		server.Health(resp, req)

		// Verify that the status code is 200 OK
		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
		}

		// Decode the actual response
		var actualResponse api.Response
		if err := json.Unmarshal(resp.Body.Bytes(), &actualResponse); err != nil {
			t.Fatalf("could not unmarshal response body: %v", err)
		}

		// Prepare expected response
		expectedResponse := api.Response{
			Data: map[string]string{"status": "pass", "version": "v0"},
		}

		// Compare the actual and expected response content
		if !compareResponses(actualResponse, expectedResponse) {
			t.Errorf("Expected response %+v, got %+v", expectedResponse, actualResponse)
		}
	})

	// Test the Health endpoint with a POST request (method not allowed)
	t.Run("Health POST request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/health", nil)
		resp := httptest.NewRecorder()

		server.Health(resp, req)

		// Verify that the status code is 405 Method Not Allowed
		if resp.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.Code)
		}
	})
}

// Helper function to compare two Response objects
func compareResponses(actual, expected api.Response) bool {
	actualJSON, _ := json.Marshal(actual)
	expectedJSON, _ := json.Marshal(expected)

	return bytes.Equal(actualJSON, expectedJSON)
}
