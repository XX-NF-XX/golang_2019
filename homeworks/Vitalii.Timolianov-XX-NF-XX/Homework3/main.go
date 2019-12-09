package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

// Products - slice of products
type Products []uint

// Machines - slice of machines that has products
type Machines []Products

func buildMachines() Machines {
	return Machines{
		Products{1, 1, 2, 3},
		Products{2, 3, 4},
		Products{1, 3},
		Products{5, 4, 3},
		Products{2, 2, 1},
	}
}

func parseFlags() Products {
	var order *[]uint = flag.UintSlice("order", nil, "Order to fulfil")
	// var rawState *[]string = flag.StringArray("state", nil, "Current state of vending machines")
	flag.Parse()
	return Products(*order)
}

// Makes shallow copy of Machines without first element at 'index'
func (machines Machines) copyWithout(index int) Machines {
	newMachines := make(Machines, len(machines))

	for i, m := range machines {
		if i == index {
			newMachines[i] = m[1:]
			continue
		}
		newMachines[i] = m
	}

	return newMachines
}

func findProduct(order Products, machines Machines) (Machines, bool) {
	if len(order) == 0 {
		return machines, true
	}
	orderedProduct := order[0]

	for i, m := range machines {
		if len(m) == 0 {
			continue
		}

		if m[0] == orderedProduct {
			reducedMachines := machines.copyWithout(i)
			orderRest := order[1:]

			usedMachines, ok := findProduct(orderRest, reducedMachines)
			if ok {
				return usedMachines, ok
			}
		}
	}

	return nil, false
}

func (machines Machines) getOrderSolutions(order Products) (Machines, bool) {
	return findProduct(order, machines)
}

func main() {
	machines := buildMachines()
	order := parseFlags()

	usedMachines, ok := machines.getOrderSolutions(order)
	if !ok {
		fmt.Printf("This order cannot be fulfilled!\n")
		return
	}

	fmt.Printf("Order fulfilled!\nState of vending machines: %v\n", usedMachines)
}
