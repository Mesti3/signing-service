package api

// TODO: REST endpoints ...
import (
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/json"
	"net/http"

	"signing-service-challenge/crypto"
	"signing-service-challenge/domain"
	"signing-service-challenge/persistence"
)

type DeviceHandler struct {
	store *persistence.InMemorys
}

var deviceRequest struct {
	Algorithm string
	Label     string
}

var privateKeyECDSA *ecdsa.PrivateKey
var publicKeyECDSA *ecdsa.PublicKey
var privateKeyRSA *rsa.PrivateKey
var publicKeyRSA *rsa.PublicKey
var device *domain.Signatured
var err error

func NewDeviceHandler(store *persistence.InMemorys) *DeviceHandler {
	return &DeviceHandler{store: store}
}

func (h *DeviceHandler) CreateDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{http.StatusText(http.StatusMethodNotAllowed)})
		return
	}

	if err := json.NewDecoder(request.Body).Decode(&deviceRequest); err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"invalid request body"})
		return
	}

	switch deviceRequest.Algorithm {
	case "RSA":
		generator := &crypto.RSAGenerator{}
		keyPair, err := generator.Generate()
		if err != nil {
			WriteInternalError(response)
			return
		}
		privateKeyRSA = keyPair.Private
		publicKeyRSA = keyPair.Public

		device = &domain.Signatured{
			Algorithm:        deviceRequest.Algorithm,
			Label:            deviceRequest.Label,
			PublicKey:        publicKeyRSA,
			PrivateKey:       privateKeyRSA,
			SignatureCounter: 0,
		}
	case "ECC":
		generator := &crypto.ECCGenerator{}
		keyPair, err := generator.Generate()
		if err != nil {
			WriteInternalError(response)
			return
		}
		privateKeyECDSA = keyPair.Private
		publicKeyECDSA = keyPair.Public

		device = &domain.Signatured{
			Algorithm:        deviceRequest.Algorithm,
			Label:            deviceRequest.Label,
			PublicKey:        publicKeyECDSA,
			PrivateKey:       privateKeyECDSA,
			SignatureCounter: 0,
		}
	default:
		WriteErrorResponse(response, http.StatusBadRequest, []string{"unsopported"})
		return
	}

	deviceID, err := h.store.CreateDevice(device)
	if err != nil {
		WriteInternalError(response)
		return
	}

	WriteAPIResponse(response, http.StatusCreated, map[string]string{
		"deviceId": deviceID,
	})
}

func (h *DeviceHandler) ListDevices(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{http.StatusText(http.StatusMethodNotAllowed)})
		return
	}

	devices := h.store.ListDevices()
	WriteAPIResponse(response, http.StatusOK, devices)
}
