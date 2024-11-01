package persistence

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"signing-service-challenge/domain"
)

type InMemorys struct {
	devices map[string]*domain.Signatured
	mut     sync.RWMutex
}

func NewInMemorys() *InMemorys {
	return &InMemorys{
		devices: make(map[string]*domain.Signatured),
	}
}

func (s *InMemorys) CreateDevice(device *domain.Signatured) (string, error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	deviceID := uuid.New().String()
	device.Id = deviceID
	s.devices[deviceID] = device

	return deviceID, nil
}

func (s *InMemorys) GetDeviceId(device *domain.Signatured) (string, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	device, exists := s.devices[device.Id]
	if !exists {
		return "", errors.New("device not found")
	}

	return device.Id, nil
}

func (s *InMemorys) ListDevices() []*domain.Signatured {
	s.mut.RLock()
	defer s.mut.RUnlock()

	devices := make([]*domain.Signatured, 0, len(s.devices))
	for _, device := range s.devices {
		devices = append(devices, device)
	}

	return devices
}
