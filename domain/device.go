package domain

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"sync"
)

type Signatured struct {
	mut              sync.RWMutex
	Id               string
	Label            string
	PublicKey        crypto.PublicKey
	PrivateKey       crypto.PrivateKey
	Algorithm        string // "RSA" or "ECC"
	SignatureCounter int
	LastSignature    string
}

func (d *Signatured) GetSignatureReference() string {
	d.mut.RLock()
	defer d.mut.RUnlock()
	if d.SignatureCounter == 0 {
		return d.Id
	}
	return d.LastSignature
}

func (d *Signatured) SignatureDevice(data string) string {
	signatureReference := d.GetSignatureReference()
	encodedSignature := base64.StdEncoding.EncodeToString([]byte(signatureReference))
	d.mut.RLock()
	securedDataToBeSigned := fmt.Sprintf("%d_%s_%s", d.SignatureCounter, data, encodedSignature)
	d.mut.RUnlock()
	return securedDataToBeSigned
}

func (d *Signatured) IncrementSignatureCounter() {
	d.mut.Lock()
	defer d.mut.Unlock()
	d.SignatureCounter++
}

func (d *Signatured) UpdateLastSignature(signature string) {
	d.mut.Lock()
	defer d.mut.Unlock()
	d.LastSignature = signature
}
