package main

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

func PrintSpin(spin [][]string) {
	colWidth := MaxSymbolWidth(spin) // compute best width for your data

	for _, row := range spin {
		for j, symbol := range row {
			if j > 0 {
				fmt.Print(" | ")
			}
			fmt.Print(PadDisplay(symbol, colWidth))
		}
		fmt.Println()
	}
}

func MaxSymbolWidth(spin [][]string) int {
	maxW := 0
	for _, row := range spin {
		for _, s := range row {
			if w := runewidth.StringWidth(s); w > maxW {
				maxW = w
			}
		}
	}
	// add a little breathing room so it doesn't look cramped
	if maxW < 2 {
		return 2
	}
	return maxW + 1
}

func PadDisplay(s string, width int) string {
	w := runewidth.StringWidth(s)
	if w >= width {
		return s
	}
	return s + strings.Repeat(" ", width-w)
}
