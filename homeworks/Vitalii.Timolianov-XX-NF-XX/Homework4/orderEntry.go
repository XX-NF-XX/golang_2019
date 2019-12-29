package main

import "sync"

// Status - Entry status
type Status int

// Enum of statuses
const (
	New Status = iota
	Pending
	Done
)

func (s Status) String() string {
	return [...]string{"new", "pending", "done"}[s]
}

// OrderEntry - Product entry in order
type OrderEntry struct {
	product  Product
	status   Status
	storages map[*Storage]bool // Storages to check
	mux      sync.Mutex
}

func newOrderEntry(products []Product) []*OrderEntry {
	entries := make([]*OrderEntry, len(products), len(products))
	for i, product := range products {
		entries[i] = &OrderEntry{
			product:  product,
			storages: make(map[*Storage]bool),
		}
	}
	return entries
}

func (e *OrderEntry) fulfill(product Product) {
	e.product = product
	e.status = Done

	for s := range e.storages {
		delete(e.storages, s)
	}
}

func (e *OrderEntry) findProduct() {
	for storage := range e.storages {
		product, ok := storage.getProduct(e.product)
		if ok {
			e.fulfill(product)
			return
		}

		delete(e.storages, storage)
	}
}

func (e *OrderEntry) checkStorages(storages ...*Storage) {
	e.mux.Lock()
	defer e.mux.Unlock()

	if e.status == Done {
		return
	}

	if e.status == New {
		e.status = Pending
	}

	for _, s := range storages {
		e.storages[s] = true
	}

	e.findProduct()
}
