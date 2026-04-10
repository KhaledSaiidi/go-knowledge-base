package lab

import "fmt"

func Hello(s string) string {
	const englishHelloPrefix = "Hello, "
	if s == "" {
		s = "World"
	}
	k := fmt.Sprintf("%s%s!", englishHelloPrefix, s)
	return k
}

func main() {
	fmt.Println(Hello("Chris"))
}
