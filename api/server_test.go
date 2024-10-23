package api_test

import (
	"net/http"
	"net/http/httptest"
	"signing-service-challenge/api"
	"testing"
)

func TestWriteInternalError(t *testing.T) {
	// Test that WriteInternalError returns a 500 Internal Server Error
	resp := httptest.NewRecorder()
	api.WriteInternalError(resp)

	if resp.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.Code)
	}

	if resp.Body.String() != http.StatusText(http.StatusInternalServerError) {
		t.Errorf("Expected response body %q, got %q", http.StatusText(http.StatusInternalServerError), resp.Body.String())
	}
}

func TestWriteErrorResponse(t *testing.T) {
	// Test WriteErrorResponse with various error cases
	resp := httptest.NewRecorder()
	errors := []string{"error1", "error2"}

	api.WriteErrorResponse(resp, http.StatusBadRequest, errors)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.Code)
	}

	expectedBody := `{"errors":["error1","error2"]}`
	if resp.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, resp.Body.String())
	}
}

func TestWriteAPIResponse(t *testing.T) {
	// Test WriteAPIResponse with a valid data struct
	resp := httptest.NewRecorder()
	data := map[string]string{"key": "value"}

	api.WriteAPIResponse(resp, http.StatusOK, data)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
	}

	expectedBody := "{\n  \"data\": {\n    \"key\": \"value\"\n  }\n}"
	if resp.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, resp.Body.String())
	}
}
