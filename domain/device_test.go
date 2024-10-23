package domain

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestSignatured_GetSignatureReference(t *testing.T) {
	// Test case where signature counter is 0, so it returns the device ID
	device := &Signatured{
		Id:               "device-123",
		SignatureCounter: 0,
		LastSignature:    "some-signature",
	}

	expected := device.Id
	result := device.GetSignatureReference()

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	// Test case where signature counter is greater than 0, so it returns the last signature
	device.SignatureCounter = 1
	expected = "some-signature"
	result = device.GetSignatureReference()

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestSignatured_SignatureDevice(t *testing.T) {
	device := &Signatured{
		Id:               "device-123",
		SignatureCounter: 0,
		LastSignature:    "",
	}

	data := "some-data"

	// Since SignatureCounter is 0, the reference should be base64(ID)
	encodedID := base64.StdEncoding.EncodeToString([]byte(device.Id))
	expected := fmt.Sprintf("%d_%s_%s", device.SignatureCounter, data, encodedID)

	result := device.SignatureDevice(data)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	// Now test with an incremented SignatureCounter
	device.SignatureCounter = 1
	device.LastSignature = "last-signature"
	encodedLastSignature := base64.StdEncoding.EncodeToString([]byte(device.LastSignature))
	expected = fmt.Sprintf("%d_%s_%s", device.SignatureCounter, data, encodedLastSignature)

	result = device.SignatureDevice(data)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestSignatured_IncrementSignatureCounter(t *testing.T) {
	device := &Signatured{
		SignatureCounter: 0,
	}

	// Call IncrementSignatureCounter
	device.IncrementSignatureCounter()

	// Check if the counter has been incremented
	if device.SignatureCounter != 1 {
		t.Errorf("expected SignatureCounter to be 1, got %d", device.SignatureCounter)
	}

	// Call again and check if it increments further
	device.IncrementSignatureCounter()

	if device.SignatureCounter != 2 {
		t.Errorf("expected SignatureCounter to be 2, got %d", device.SignatureCounter)
	}
}

func TestSignatured_UpdateLastSignature(t *testing.T) {
	device := &Signatured{
		LastSignature: "old-signature",
	}

	// Update the last signature
	newSignature := "new-signature"
	device.UpdateLastSignature(newSignature)

	// Check if the last signature is updated
	if device.LastSignature != newSignature {
		t.Errorf("expected LastSignature to be %s, got %s", newSignature, device.LastSignature)
	}
}
