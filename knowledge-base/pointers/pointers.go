package main

import "fmt"

// p := &i -> & is a pointer (an address of a variable)
// *p -> * is dereferencing, it gives the value of the address
// p *int -> is a pointer to an int, it can STORE an address of an int
// *T works for all types, including structs
//----------------------------------------------------------------------------------------------------------------------------
// slice, map, channel, func share underlying data (reference-like, not real pointers)
// passing them copies the variable, but both still point to the same data → changes are visible outside
//----------------------------------------------------------------------------------------------------------------------------

func pointer() {
	i, j := 42, 2701
	p := &i // -> & is address of i
	fmt.Println("Print the address of i:", p)
	fmt.Println("Print the value of i:", *p) // -> * is dereferencing, it gives the value of the address
	*p = 21
	fmt.Println("Print the value of i after modification:", i)

	p = &j
	*p = *p / 37
	fmt.Println("Print the value of j after modification:", j)
}

func squareAdd(p *int) {
	*p *= *p
	fmt.Printf("Print the value of the squared number: %d with address %p\n", *p, p)
}

type person struct {
	name string
	age  int
}

func initPerson() *person {
	p := person{"Alice", 30}
	return &p
}
