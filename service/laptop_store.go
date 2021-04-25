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
	Find(string) (*pb.Laptop, error)
	Search(*pb.Filter, func(*pb.Laptop) error) error
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
	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	s.data[other.Id] = other
	return nil
}

func (s *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	laptop := s.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	return deepCopy(laptop)
}

func toBit(memory *pb.Memory) uint64 {
	value := memory.GetValue()

	switch memory.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3 // 8 = 2^3
	case pb.Memory_KILOBYTE:
		return value << 13 // 1024 *8 = 2* 10 * 2^3 = 2* 13
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 44
	default:
		return 0
	}
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}

	if laptop.GetCpu().GetNumberCores() > filter.GetMinCpuCores() {
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}

	if toBit(laptop.GetMemory()) < toBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func (s *InMemoryLaptopStore) Search(filter *pb.Filter, found func(laptop *pb.Laptop) error) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, laptop := range s.data {
		if isQualified(filter, laptop) {
			// deep cory
			other, err := deepCopy(laptop)
			if err != nil {
				return err
			}

			if err := found(other); err != nil {
				return err
			}
		}
	}

	return nil
}

func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	other := &pb.Laptop{}
	if err := copier.Copy(other, laptop); err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}
