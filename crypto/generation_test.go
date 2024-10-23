package crypto

import (
	"testing"
)

func TestRSAGenerator_Generate(t *testing.T) {
	generator := RSAGenerator{}

	// Test RSA key pair generation
	keyPair, err := generator.Generate()
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Check that the key pair is not nil
	if keyPair == nil {
		t.Fatal("Generated RSA key pair is nil")
	}

	// Check that both the public and private keys are not nil
	if keyPair.Private == nil {
		t.Error("Generated RSA private key is nil")
	}

	if keyPair.Public == nil {
		t.Error("Generated RSA public key is nil")
	}
}

func TestECCGenerator_Generate(t *testing.T) {
	generator := ECCGenerator{}

	// Test ECC key pair generation
	keyPair, err := generator.Generate()
	if err != nil {
		t.Fatalf("Failed to generate ECC key pair: %v", err)
	}

	// Check that the key pair is not nil
	if keyPair == nil {
		t.Fatal("Generated ECC key pair is nil")
	}

	// Check that both the public and private keys are not nil
	if keyPair.Private == nil {
		t.Error("Generated ECC private key is nil")
	}

	if keyPair.Public == nil {
		t.Error("Generated ECC public key is nil")
	}
}
