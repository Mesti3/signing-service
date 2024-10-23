package persistence

import (
	"testing"

	"signing-service-challenge/domain"
)

func TestInMemorys_CreateDevice(t *testing.T) {
	store := NewInMemorys()

	// Create a dummy device
	device := &domain.Signatured{
		Label:            "Test Device",
		SignatureCounter: 0,
	}

	// Test device creation
	deviceID, err := store.CreateDevice(device)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deviceID == "" {
		t.Fatal("expected a valid device ID, got an empty string")
	}

	// Verify that the device is stored
	storedDevice, exists := store.devices[deviceID]
	if !exists {
		t.Fatalf("expected device to exist in store")
	}

	if storedDevice.Label != device.Label {
		t.Errorf("expected device label %v, got %v", device.Label, storedDevice.Label)
	}
}

func TestInMemorys_GetDeviceId(t *testing.T) {
	store := NewInMemorys()

	// Create a dummy device
	device := &domain.Signatured{
		Label:            "Test Device",
		SignatureCounter: 0,
	}

	// Create device
	deviceID, err := store.CreateDevice(device)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Retrieve device ID
	retrievedID, err := store.GetDeviceId(device)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if retrievedID != deviceID {
		t.Errorf("expected device ID %v, got %v", deviceID, retrievedID)
	}
}

func TestInMemorys_GetDeviceId_DeviceNotFound(t *testing.T) {
	store := NewInMemorys()

	// Try retrieving an ID from a non-existent device
	device := &domain.Signatured{Id: "non-existent-id"}
	_, err := store.GetDeviceId(device)
	if err == nil {
		t.Fatal("expected an error, got none")
	}

	expectedErr := "device not found"
	if err.Error() != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err.Error())
	}
}

func TestInMemorys_ListDevices(t *testing.T) {
	store := NewInMemorys()

	// Create dummy devices
	device1 := &domain.Signatured{
		Label:            "Device 1",
		SignatureCounter: 0,
	}
	device2 := &domain.Signatured{
		Label:            "Device 2",
		SignatureCounter: 0,
	}

	// Add devices to the store
	store.CreateDevice(device1)
	store.CreateDevice(device2)

	// List devices
	devices := store.ListDevices()

	if len(devices) != 2 {
		t.Fatalf("expected 2 devices, got %d", len(devices))
	}

	expectedLabels := []string{"Device 1", "Device 2"}
	for i, device := range devices {
		if device.Label != expectedLabels[i] {
			t.Errorf("expected device label %v, got %v", expectedLabels[i], device.Label)
		}
	}
}
