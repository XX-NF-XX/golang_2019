package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	MapVsSlice()
	MapWithCapacity()
}

// MapVsSlice - compare time of find random element in slice vs map
func MapVsSlice() {
	lookup := map[string]int{
		"one":   0,
		"two":   1,
		"three": 2,
		"four":  3,
		"five":  4,
		"six":   5,
		"seven": 6,
		"eight": 7,
		"nine":  8,
		"ten":   9,
	}
	values := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	search := []string{"one", "three", "five", "six", "eight", "ten"}
	temp := 0
	t0 := time.Now()

	// Version 1: search map with lookup.
	for i := 0; i < 10000000; i++ {
		needle := search[rand.Intn(len(search)-1)]
		if v, ok := lookup[needle]; ok {
			temp = v
		}
	}

	t1 := time.Now()

	// Version 2: search slice with for-loop.
	for i := 0; i < 10000000; i++ {
		for x := range values {
			needle := search[rand.Intn(len(search)-1)]
			if values[x] == needle {
				temp = x
				break
			}
		}
	}

	t2 := time.Now()
	// Benchmark results.
	fmt.Println(temp)
	fmt.Println("Map lookup  : ", t1.Sub(t0))
	fmt.Println("Slice search: ", t2.Sub(t1))
}

// MapWithCapacity - compare speed of adding values to maps with predefined capacity and not
func MapWithCapacity() {
	t0 := time.Now()

	// Version 1: use map with exact capacity.
	for i := 0; i < 10000; i++ {
		values := make(map[int]int, 1000)
		for x := 0; x < 1000; x++ {
			values[x] = x
		}
		if values[0] != 0 {
			fmt.Println(0)
		}
	}

	t1 := time.Now()

	// Version 2: no capacity.
	for i := 0; i < 10000; i++ {
		values := map[int]int{}
		for x := 0; x < 1000; x++ {
			values[x] = x
		}
		if values[0] != 0 {
			fmt.Println(0)
		}
	}

	t2 := time.Now()
	// Benchmark results.
	fmt.Println("With predefined capacity   : ", t1.Sub(t0))
	fmt.Println("Without predefined capacity: ", t2.Sub(t1))
}
