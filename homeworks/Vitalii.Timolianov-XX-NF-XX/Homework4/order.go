package main

import (
	"fmt"
	"sync"

	"github.com/rs/xid"
)

// Order - order to fulfill
type Order struct {
	entries []*OrderEntry
	id      string
	mux     sync.Mutex
}

func newOrder(products []Product) Order {
	return Order{
		entries: newOrderEntry(products),
		id:      xid.New().String(),
	}
}

func (o *Order) checkStorage(s *Storage) {
	o.mux.Lock()
	defer o.mux.Unlock()

	for _, entry := range o.entries {
		go entry.checkStorages(s)
	}
}

func (o *Order) getStatus() Status {
	if len(o.entries) < 1 {
		return Done
	}

	status := o.entries[0].status // use as default status

	for _, p := range o.entries {
		switch p.status {
		case Pending:
			return Pending // if any product has Pending status
		case New, Done:
			if status != p.status {
				return Pending // if products has different statuses
			}
		default:
			fmt.Printf("Unknown status %v\n", p.status.String())
		}
	}

	return status // this happens only when all products have the same status (new or done)
}

func (o *Order) cancel() (products []Product) {
	products = []Product{}

	for _, entry := range o.entries {
		if entry.status == Done {
			products = append(products, entry.product)
		}
	}
	return
}
