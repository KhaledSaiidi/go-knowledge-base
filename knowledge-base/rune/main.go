package main

import "fmt"

func main() {
	a := "S"
	b := 'S'
	c := string(b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	log := fmt.Sprintf("a: %T, b: %T, c: %T", a, b, c)
	fmt.Println(log)
	for _, r := range log {
		fmt.Print(string(r) + " ")
	}
	fmt.Println()

	s := "Hello !"
	fmt.Println(len(s)) // -> 7 : each character is 1 byte
	r := "hello ❗"
	fmt.Println(len(r)) // -> 9 : each character is 1 byte + ❗ is Unicode: U+2757 In UTF-8, it is encoded using 3 bytes

}
