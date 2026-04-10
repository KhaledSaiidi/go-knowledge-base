package main

import (
	"math/rand"
)

func SpinReels(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func getRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

func GetSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}
	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}
	for c := 0; c < cols; c++ {
		selected := map[int]bool{}
		for r := 0; r < rows; r++ {
			for true {
				reandomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[reandomIndex]
				if !exists {
					selected[reandomIndex] = true
					result[r] = append(result[r], reel[reandomIndex])
					break
				}
			}
		}
	}
	return result
}
