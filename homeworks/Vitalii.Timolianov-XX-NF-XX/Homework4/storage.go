package main

import (
	"sync"
	"fmt"
)

// Storage - product container
type Storage struct {
	products        []Product
	updatedStorages chan *Storage
	mut             sync.Mutex
}

func newStorage(products []Product, updatedStorages chan *Storage) *Storage {
	return &Storage{
		products:        products,
		updatedStorages: updatedStorages,
	}
}

func (s *Storage) getProduct(p Product) (product Product, ok bool) {
	s.mut.Lock()
	defer s.mut.Unlock()

	if len(s.products) < 1 {
		return
	}

	if s.products[0] == p {
		product, ok = s.products[0], true
		s.products = append(s.products[1:])

		if s.updatedStorages != nil {
			s.updatedStorages <- s
		}
	}

	return
}

func (s *Storage) addProduct(products ...Product) {
	s.products = append(s.products, products...)
}

func (s *Storage) String() string {
	return fmt.Sprintf("%v", s.products)
}
