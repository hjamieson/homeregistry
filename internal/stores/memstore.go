package stores

import (
	"maps"
	"sync"

	"github.com/hjamieson/homeregistry/internal/model"
)

type MemStore struct {
	lock sync.Mutex
	mn   map[string]model.Device
}

func NewMemStore() *MemStore {
	return &MemStore{
		lock: sync.Mutex{},
		mn:   make(map[string]model.Device),
	}
}

func (s *MemStore) Put(d model.Device) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.mn[d.Id] = d
	return nil
}

func (s *MemStore) Get(id string) (*model.Device, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	dev, ok := s.mn[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &dev, nil
}

func (s *MemStore) List() ([]model.Device, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	result := make([]model.Device, 0, len(s.mn))
	for dev := range maps.Values(s.mn) {
		result = append(result, dev)
	}
	return result, nil
}
