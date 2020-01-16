package main

import "fmt"

// Store - unit that fulfills orders using its storages
type Store struct {
	orders          map[string]*Order
	storages        []*Storage
	updatedStorages chan *Storage
	canceled        *Storage
}

func defaultStorages(updatedStorages chan *Storage) []*Storage {
	storages := []*Storage{
		newStorage([]Product{1, 1, 2, 3}, updatedStorages),
		newStorage([]Product{2, 3, 4}, updatedStorages),
		newStorage([]Product{1, 3}, updatedStorages),
		newStorage([]Product{5, 4, 3}, updatedStorages),
		newStorage([]Product{2, 2, 1}, updatedStorages),
	}

	fmt.Println("Created default storages:");
	for i, storage := range storages {
		fmt.Printf("storage #%v: %v\n", i, storage)
	}

	return storages;
}

func (s *Store) addUpdatedStorages() {
	for storage := range s.updatedStorages {
		for _, order := range s.orders {
			order.checkStorage(storage)
		}
	}
}

func newStore() Store {
	updatedStorages := make(chan *Storage)

	store := Store{
		orders:          make(map[string]*Order),
		updatedStorages: updatedStorages,
		storages:        defaultStorages(updatedStorages),
		canceled:        newStorage([]Product{}, nil),
	}

	go store.addUpdatedStorages()

	return store
}

func (s *Store) createOrder(products []Product) string {
	s.log()
	order := newOrder(products)
	s.orders[order.id] = &order

	for _, entry := range order.entries {
		go entry.checkStorages(s.storages...)
	}

	return order.id
}

func (s *Store) getOrder(id string) (o *Order, ok bool) {
	s.log()
	o, ok = s.orders[id]
	return
}

func (s *Store) deleteOrder(id string) (o *Order, ok bool) {
	o, ok = s.getOrder(id)

	if ok {
		products := o.cancel()
		s.canceled.addProduct(products...)
		delete(s.orders, id)
	}

	s.log()
	return
}

func (s *Store) log() {
	fmt.Printf("Storages:\n")
	for i, storage := range s.storages {
		fmt.Printf("%v: %v\n", i, storage)
	}
	fmt.Printf("Canceled: %v\n", s.canceled)
}
