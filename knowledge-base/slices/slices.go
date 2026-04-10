package main

import (
	"fmt"
	"slices"
)

func reminder() {
	var s1 []string = []string{"a", "b", "c"}
	var s2 []string = []string{"d", "e", "f"}
	isEqual := slices.Equal(s1, s2)
	contains := slices.Contains(s2, "d")
	s3 := append(s1, "d")
	s4 := append(s1, s2...)
	s5 := make([]string, len(s4), 6)
	copy(s5, s4)
	fmt.Println("s1 and s2 are equal:", isEqual)
	fmt.Println("s2 contains 'd':", contains)
	fmt.Println("s3 (s1 with 'd'):", s3)
	fmt.Println("s4 (s1 with s2):", s4)
	fmt.Println("s5 (copy of s4):", s5)
}

var users = []string{}

func addUsers(user ...string) {
	users = append(users, user...)
	fmt.Println("Users:", users)
}

func removeUsers(index int) error {
	if index < 0 || index >= len(users) {
		return fmt.Errorf("index out of range: %d", index)
	}
	users := append(users[:index], users[index+1:]...)
	fmt.Println("Users:", users)
	return nil
}

type Myslice []int

// numbers := Myslice{1, 2, 3, 4, 5}
func (m Myslice) Remove(index int) ([]int, error) {
	if index < 0 || index >= len(m) {
		return m, fmt.Errorf("index out of range: %d", index)
	}
	// `append` can reuse the same backing array, so removing this way can also mutate the original slice view.
	s := append(m[:index], m[index+1:]...)
	fmt.Println("Initslice:", m)
	fmt.Println("Myslice:", s)
	return s, nil
}

// Sorting:
type SortNumbers []int

func (n SortNumbers) Len() int           { return len(n) }
func (n SortNumbers) Less(i, j int) bool { return n[i] < n[j] }
func (n SortNumbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

type byInc struct {
	SortNumbers
}

func (n byInc) Len() int           { return len(n.SortNumbers) }
func (n byInc) Less(i, j int) bool { return n.SortNumbers[i] < n.SortNumbers[j] }
func (n byInc) Swap(i, j int) {
	n.SortNumbers[i], n.SortNumbers[j] = n.SortNumbers[j], n.SortNumbers[i]
}

type byDec struct {
	SortNumbers
}

func (n byDec) Len() int           { return len(n.SortNumbers) }
func (n byDec) Less(i, j int) bool { return n.SortNumbers[i] > n.SortNumbers[j] }
func (n byDec) Swap(i, j int) {
	n.SortNumbers[i], n.SortNumbers[j] = n.SortNumbers[j], n.SortNumbers[i]
}
