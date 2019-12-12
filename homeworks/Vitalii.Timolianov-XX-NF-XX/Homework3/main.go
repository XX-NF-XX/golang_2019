package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	flag "github.com/spf13/pflag"
)

var defaultOrder = []uint{1, 2, 1, 3, 4}
var defaultState = "[[1,1,2,3],[2,3,4],[1,3],[5,4,3],[2,2,1]]"

// Products - slice of products
type Products []uint

// Machines - slice of machines with products
type Machines []Products

// Converts JSON state string to Machines
func getMachinesFromState(state *string) Machines {
	var stateBytes = []byte(*state)
	var machines Machines

	err := json.Unmarshal(stateBytes, &machines)

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	return machines
}

// Converts parameters to valid data
func parseFlags() (Products, Machines) {
	var order *[]uint = flag.UintSlice("order", defaultOrder, "Order to fulfill e.g.: {4,1}")
	var rawState *string = flag.String("state", defaultState, "Current state of vending machines e.g.: [[1,2],[4,3]]")

	flag.Parse()

	machines := getMachinesFromState(rawState)
	products := Products(*order)

	return products, machines
}

// Closes channel when wait group is done (I don't like it, but it works)
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

// Finds the top order product in machines
// All the logic/magic happens here
func findProduct(order Products, machines Machines, c chan Machines, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	orderedProduct := order[0]
	for i, m := range machines {
		if len(m) == 0 {
			continue
		}

		if m[0] == orderedProduct {
			reducedMachines := machines.copyWithout(i)
			orderRest := order[1:]

			if len(orderRest) == 0 {
				c <- reducedMachines
				return
			}

			waitgroup.Add(1)
			go findProduct(orderRest, reducedMachines, c, waitgroup)
		}
	}
}

// Tries to resolve order and return possible states of vending machines
func (machines Machines) getOrderSolutions(order Products) ([]Machines, bool) {
	var waitgroup sync.WaitGroup
	var possibleStates = []Machines{}

	c := make(chan Machines, 8)

	waitgroup.Add(1)
	go findProduct(order, machines, c, &waitgroup)
	go closeOnDone(c, &waitgroup)

	for m := range c {
		possibleStates = append(possibleStates, m)
	}

	ok := len(possibleStates) > 0
	return possibleStates, ok
}

// Yep, this is where the whole thing starts
func main() {
	order, machines := parseFlags()
	fmt.Printf("order: %v\n", order)
	fmt.Printf("state: %v\n\n", machines)

	possibleStates, ok := machines.getOrderSolutions(order)
	if !ok {
		fmt.Printf("This order cannot be fulfilled!\n")
		os.Exit(1)
	}

	fmt.Printf("Order can be fulfilled in %v way(s).\n", len(possibleStates))
	fmt.Printf("\nPossible vending machine states after fulfillment:\n")

	for i, s := range possibleStates {
		fmt.Printf("#%v %v\n", i+1, s)
	}
}
