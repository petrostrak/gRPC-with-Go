package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/petrostrak/gRPC-with-Go/pb"
)

var (
	ErrAlreadyExists = errors.New("record already exists")
)

type LaptopStore interface {
	Save(*pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (s *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &pb.Laptop{}
	if err := copier.Copy(other, laptop); err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	s.data[other.Id] = other
	return nil
}
