package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/asn1"
	"errors"
	"math/big"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

type RSASigner struct {
	privateKey *rsa.PrivateKey
}

type ECDSASigner struct {
	privateKey *ecdsa.PrivateKey
}

func NewRSASigner(privateKey *rsa.PrivateKey) *RSASigner {
	return &RSASigner{privateKey: privateKey}
}

func (s *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	hashed := sha256.Sum256(dataToBeSigned)
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func NewECDSASigner(privateKey *ecdsa.PrivateKey) *ECDSASigner {
	return &ECDSASigner{privateKey: privateKey}
}

func (s *ECDSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	hashed := sha256.Sum256(dataToBeSigned)
	r, sVal, err := ecdsa.Sign(rand.Reader, s.privateKey, hashed[:])
	if err != nil {
		return nil, err
	}
	signature, err := asn1.Marshal(struct{ R, S *big.Int }{R: r, S: sVal})
	if err != nil {
		return nil, errors.New("failed to marshal ecdsa signature")
	}

	return signature, nil
}
