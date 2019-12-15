package main

import "github.com/rs/xid"

// Products - list of products
type Products []int

// Status - Order status
type Status int

// Enum of order statuses
const (
	New Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	return [...]string{"new", "in progress", "done"}[s]
}

// Order - order to fulfill
type Order struct {
	Products Products `json:"products"`
	ID       string   `json:"id"`
	Status   Status   `json:"status"`
}

// Store - unit that fulfills orders within its storages
type Store struct {
	Orders   map[string]*Order
	Storages []Products
}

func newStore() Store {
	return Store{
		Orders:   make(map[string]*Order),
		Storages: make([]Products, 0),
	}
}

func (s *Store) addOrder(o *Order) {
	o.ID = xid.New().String()
	s.Orders[o.ID] = o
}

func (s *Store) getOrder(id string) (o *Order, ok bool) {
	o, ok = s.Orders[id]
	return
}

func (s *Store) deleteOrder(id string) (o *Order, ok bool) {
	o, ok = s.getOrder(id)
	delete(s.Orders, id)
	return
}
