package main

import "fmt"

type MP map[string]int

func (m MP) Get(key string) (int, bool) {
	value, ok := m[key]
	return value, ok
}
func (m MP) Set(key string, value int) {
	m[key] = value
}
func (m MP) Delete(key string) {
	delete(m, key)
}
func (m MP) String() string {
	result := "{"
	for key, value := range m {
		x := fmt.Sprintf("%d", value)
		result += key + ": " + x + ", "
	}
	if len(result) > 1 {
		result = result[:len(result)-2] // Remove the trailing comma and space,
	}
	result += "}"
	return result
}
