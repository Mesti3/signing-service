package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"signing-service-challenge/api"
	"signing-service-challenge/domain"
	"signing-service-challenge/persistence"
	"strings"
	"testing"
)

type CreateDeviceResponse struct {
	Data struct {
		DeviceID string `json:"deviceId"`
	} `json:"data"`
}

// Helper function to write JSON body for the request
func createDeviceRequestBody(algorithm, label string) *bytes.Buffer {
	body := map[string]string{
		"Algorithm": algorithm,
		"Label":     label,
	}
	bodyBytes, _ := json.Marshal(body)
	return bytes.NewBuffer(bodyBytes)
}

func TestDeviceHandler_CreateDevice(t *testing.T) {
	// Set up the in-memory store
	store := persistence.NewInMemorys()
	deviceHandler := api.NewDeviceHandler(store)

	// Create a request to create a new device
	reqBody := `{"algorithm": "RSA", "label": "Test Device"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v0/devices", bytes.NewBufferString(reqBody))
	resp := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(deviceHandler.CreateDevice)
	handler.ServeHTTP(resp, req)

	// Check the response status code
	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.Code)
	}
	t.Logf("Response body: %s", resp.Body.Bytes())

	// Decode the response into a map
	var response CreateDeviceResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check that the 'deviceId' is present in the response
	deviceId := response.Data.DeviceID
	if deviceId == "" {
		t.Error("Expected 'deviceId' in the response")
	} else {
		t.Logf("Device ID created: %s", deviceId)
	}
}

func TestDeviceHandler_ListDevices(t *testing.T) {
	store := persistence.NewInMemorys()
	handler := api.NewDeviceHandler(store)

	// Simulate device creation
	store.CreateDevice(&domain.Signatured{
		Algorithm:        "RSA",
		Label:            "TestDevice",
		PublicKey:        nil,
		PrivateKey:       nil,
		SignatureCounter: 0,
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v0/devices/list", nil)
	resp := httptest.NewRecorder()

	handler.ListDevices(resp, req)

	// Verify response status code
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
	}

	// Verify that the response contains the list of devices
	body := resp.Body.String()
	if !strings.Contains(body, "TestDevice") {
		t.Error("Expected 'TestDevice' in the response")
	}
}

func TestDeviceHandler_CreateDevice_MethodNotAllowed(t *testing.T) {
	store := persistence.NewInMemorys()
	handler := api.NewDeviceHandler(store)

	// Create a GET request which is not allowed for CreateDevice
	req := httptest.NewRequest(http.MethodGet, "/api/v0/devices", nil)
	resp := httptest.NewRecorder()

	handler.CreateDevice(resp, req)

	// Verify that the response status code is 405 Method Not Allowed
	if resp.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.Code)
	}
}

func TestDeviceHandler_ListDevices_MethodNotAllowed(t *testing.T) {
	store := persistence.NewInMemorys()
	handler := api.NewDeviceHandler(store)

	// Create a POST request which is not allowed for ListDevices
	req := httptest.NewRequest(http.MethodPost, "/api/v0/devices/list", nil)
	resp := httptest.NewRecorder()

	handler.ListDevices(resp, req)

	// Verify that the response status code is 405 Method Not Allowed
	if resp.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.Code)
	}
}
