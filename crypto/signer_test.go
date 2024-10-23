package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"testing"
)

func TestRSASigner_Sign(t *testing.T) {
	// Generate a test RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key: %v", err)
	}

	// Create a new RSASigner with the generated private key
	signer := NewRSASigner(privateKey)

	// Test data to be signed
	data := []byte("test data for RSA signing")

	// Call the Sign method
	signature, err := signer.Sign(data)
	if err != nil {
		t.Fatalf("RSASigner.Sign() failed: %v", err)
	}

	// Verify that the signature is not empty
	if len(signature) == 0 {
		t.Error("RSASigner.Sign() returned an empty signature")
	}

	// Optionally, verify the signature using rsa.VerifyPKCS1v15 to check its validity
	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		t.Errorf("Failed to verify RSA signature: %v", err)
	}
}

func TestECDSASigner_Sign(t *testing.T) {
	// Generate a test ECDSA private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate ECDSA private key: %v", err)
	}

	// Create a new ECDSASigner with the generated private key
	signer := NewECDSASigner(privateKey)

	// Test data to be signed
	data := []byte("test data for ECDSA signing")

	// Call the Sign method
	signature, err := signer.Sign(data)
	if err != nil {
		t.Fatalf("ECDSASigner.Sign() failed: %v", err)
	}

	// Verify that the signature is not empty
	if len(signature) == 0 {
		t.Error("ECDSASigner.Sign() returned an empty signature")
	}
}
