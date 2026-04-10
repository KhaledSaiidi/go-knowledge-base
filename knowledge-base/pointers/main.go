package main

import "fmt"

func main() {
	pointer()
	println("---------------------------------")
	a := 4
	squareAdd(&a)
	println("---------------------------------")
	fmt.Println("Print the person struct:", *initPerson())
}
