package main

import (
	"fmt"
	"sync"

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

func closeOnDone(c chan Machines, waitgroup *sync.WaitGroup) {
	waitgroup.Wait()
	close(c)
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

func findProduct(order Products, machines Machines, c chan Machines, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	if len(order) == 0 {
		c <- machines
		return
	}

	orderedProduct := order[0]
	for i, m := range machines {
		if len(m) == 0 {
			continue
		}

		if m[0] == orderedProduct {
			reducedMachines := machines.copyWithout(i)
			orderRest := order[1:]

			waitgroup.Add(1)
			go findProduct(orderRest, reducedMachines, c, waitgroup)
		}
	}
}

func (machines Machines) getOrderSolutions(order Products) ([]Machines, bool) {
	var waitgroup sync.WaitGroup
	c := make(chan Machines, 4)
	var possibleStates = []Machines{}

	waitgroup.Add(1)
	go findProduct(order, machines, c, &waitgroup)
	go closeOnDone(c, &waitgroup)

	for m := range c {
		possibleStates = append(possibleStates, m)
	}

	if len(possibleStates) <= 0 {
		return nil, false
	}
	return possibleStates, true
}

func main() {
	machines := buildMachines()
	order := parseFlags()

	usedMachines, ok := machines.getOrderSolutions(order)
	if !ok {
		fmt.Printf("This order cannot be fulfilled!\n")
		return
	}

	fmt.Printf("Order can be fulfilled in %v different ways.\n", len(usedMachines))
	fmt.Printf("\nPossible states:\n")
	for i, m := range usedMachines {
		fmt.Printf("#%v %v\n", i+1, m)
	}
}
