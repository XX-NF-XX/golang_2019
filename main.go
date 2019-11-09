package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {

	//arrays
	fmt.Printf("\nArrays:\n")

	var arr [5]int
	fmt.Printf(" arr: (%T) %d %d %v\n", arr, len(arr), cap(arr), arr)

	arr1 := [6]int{1, 2, 3, 4, 5}
	fmt.Printf("arr1: (%T) %d %d %v\n", arr1, len(arr1), cap(arr1), arr1)

	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr2: (%T) %d %d %v\n", arr2, len(arr2), cap(arr2), arr2)

	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("arr3: (%T) %p %d %d %v\n", arr3, arr3, len(arr3), cap(arr3), arr3)

	//slices

	fmt.Printf("\nSlices:\n")

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf(" a: (%T)  %d %d %v\n", a, len(a), cap(a), a)
	a1 := a[3:7]
	fmt.Printf("[3:7]: (%T) %p %d %d %v\n", a1, a1, len(a1), cap(a1), a1)
	a2 := a[:5]
	fmt.Printf("[ :5]: (%T) %p %d %d %v\n", a2, a2, len(a2), cap(a2), a2)
	a3 := a[4:]
	fmt.Printf("[4: ]: (%T) %p %d %d %v\n", a3, a3, len(a3), cap(a3), a3)

	fmt.Printf("\nAfter append 10 to alice a1\n")
	a1 = append(a1, 10)
	a3 = append(a3, 11)

	fmt.Printf(" a: (%T)  %d %d %v\n", a, len(a), cap(a), a)
	fmt.Printf("[3:7]: (%T) %p %d %d %v\n", a1, a1, len(a1), cap(a1), a1)
	fmt.Printf("[ :5]: (%T) %p %d %d %v\n", a2, a2, len(a2), cap(a2), a2)
	fmt.Printf("[4: ]: (%T) %p %d %d %v\n", a3, a3, len(a3), cap(a3), a3)

	fmt.Printf("\nStrings as arrays\n")

	str := "some string with digits 345 and other symbols"
	b := []byte(str)
	fmt.Printf(" b: (%T) %p %d %d [%s]\n", b, b, len(b), cap(b), b)

	re, _ := regexp.Compile(`[0-9]+`)

	s := re.Find(b)
	fmt.Printf(" s: (%T) %p %d %d [%s]\n", s, s, len(s), cap(s), s)

	s[0] = 'A'
	fmt.Printf(" b: (%T) %p %d %d [%s]\n", b, b, len(b), cap(b), b)
	fmt.Printf(" s: (%T) %p %d %d [%s]\n", s, s, len(s), cap(s), s)

	s1 := make([]byte, len(s))
	copy(s1, s)
	fmt.Printf("s1: (%T) %p %d %d [%s]\n", s1, s1, len(s1), cap(s1), s1)

	s[0] = 'B'
	s1[0] = 'X'
	fmt.Printf(" b: (%T) %p %d %d [%s]\n", b, b, len(b), cap(b), b)
	fmt.Printf(" s: (%T) %p %d %d [%s]\n", s, s, len(s), cap(s), s)
	fmt.Printf("s1: (%T) %p %d %d [%s]\n", s1, s1, len(s1), cap(s1), s1)

	fmt.Printf("\n\n\nMake slices:\n")

	ns1 := make([]int, 5, 5)
	fmt.Printf("make slice 1: (%T) %p %d %d %v\n", ns1, ns1, len(ns1), cap(ns1), ns1)

	ns2 := make([]int, 3, 5)
	fmt.Printf("make slice 2: (%T) %p %d %d %v\n", ns2, ns2, len(ns2), cap(ns2), ns2)

	ap := make([]int, 5, 5)
	fmt.Printf("ap: (%T) %p %d %d %v\n", ap, ap, len(ap), cap(ap), ap)
	for i := 0; i < 20; i++ {
		ap = append(ap, i)
		fmt.Printf("ap: (%T) %p %d %d %v\n", ap, ap, len(ap), cap(ap), ap)
	}
	fmt.Printf("\n\n")

	ap1 := make([]int, 5, 5)
	fmt.Printf("ap1: (%T) %p %d %d\n", ap1, ap1, len(ap1), cap(ap1))
	for i := 0; i < 2000; i++ {
		ap1 = append(ap1, i)
		if i%100 == 0 {
			fmt.Printf("ap1: (%T) %p %d %d\n", ap1, ap1, len(ap1), cap(ap1))
		}
	}

	//maps
	fmt.Printf("\nMaps:\n")

	m := map[string]string{
		"lemon":   "yellow",
		"apple":   "red",
		"avocado": "green",
	}

	fmt.Printf("m: (%T) %p %d %v\n", m, m, len(m), m)
	//
	appleColor := m["apple"]
	//
	fmt.Printf("\nApple is '%s'\n", appleColor)
	//
	bananaColor := m["babana"]
	fmt.Printf("\nBanana is '%s'\n", bananaColor)
	// //
	pineAppleColor, ok := m["pineapple"]
	if ok {
		fmt.Printf("\nPineapple is '%s'\n", pineAppleColor)
	} else {
		fmt.Printf("\nPineapple not exist\n")
	}
	// // //
	m["pineapple"] = "browne"
	fmt.Printf("m: (%T) %p %d %v\n", m, m, len(m), m)

	delete(m, "avocado")
	fmt.Printf("m: (%T) %p %d %v\n", m, m, len(m), m)
	// // //
	fmt.Printf("\n")
	for key, value := range m {
		fmt.Printf("%s: %s\n", key, value)
	}
	fmt.Printf("\n")
	//
	//get map keys slice
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}

	fmt.Printf("Keys: %v\n\n\n", keys)

	// show that keys in map not ordered
	// // Loop ten times.
	for i := 0; i < 10; i++ {
		// Print all keys in range loop over map.
		// ... Ordering is randomized.
		for key := range m {
			fmt.Print(key)
			fmt.Print(" ")
		}
		fmt.Println()
	}

	//
	// sorting using sort package
	//
	fmt.Printf("\nKeys        : %v\n", keys)
	sort.Strings(keys)
	fmt.Printf("Sorted Keys : %v\n", keys)

	sort.Sort(SSlice(keys))
	fmt.Printf("Reverse Keys: %v\n", keys)

	a := []string{}
	b := SSlice{}

	a = b

	fmt.Printf("Variable a: %v\n", a)

}
