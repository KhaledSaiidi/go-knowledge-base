package main

func main() {
	myMap := MP{"five": 5, "two": 2, "nine": 9, "one": 1, "six": 6}
	value, ok := myMap.Get("two")
	if ok {
		println("Value for 'two':", value)
	} else {
		println("'two' not found in the map")
	}

	myMap.Set("three", 3)
	value, ok = myMap.Get("three")
	if ok {
		println("Value for 'three':", value)
		println(myMap.String())
	} else {
		println("'three' not found in the map")
	}
	myMap.Delete("five")
	value, ok = myMap.Get("five")
	if ok {
		println("Value for 'five':", value)
		println(myMap.String())
	} else {
		println("'five' not found in the map")
	}
}
