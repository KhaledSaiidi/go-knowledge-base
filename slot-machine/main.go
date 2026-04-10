package main

import (
	"fmt"
)

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	symbols := map[string]uint{
		" 7": 3,
		"💎":  5,
		"🍒":  6,
		"★":  8,
	}

	multipliers := map[string]uint{
		"7": 20,
		"💎": 8,
		"🍒": 5,
		"★": 2,
	}

	symbolArr := SpinReels(symbols)
	// fmt.Println(symbolArr)

	balance := uint(50)
	GetName()
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			fmt.Println("Thank you for playing at KS Casino. Goodbye!")
			break
		}
		balance -= bet
		fmt.Printf("You have bet $%d. Good luck!\n", bet)
		spinResult := GetSpin(symbolArr, 4, 3)
		PrintSpin(spinResult)
		winngLines := checkWin(spinResult, multipliers)

		for i, multi := range winngLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("You won $%d, (%dx) on line #%d!\n", win, multi, i+1)
			}
		}
	}
	fmt.Printf("You left with $%d. To next time!\n", balance)

}
