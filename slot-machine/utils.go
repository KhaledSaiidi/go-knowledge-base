package main

import "fmt"

func GetName() string {
	name := ""
	fmt.Println("Welcome to KS Casino...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("You have a balance of $%d. Enter your bet: ", balance)
		_, err := fmt.Scanln(&bet)
		if err != nil {
			fmt.Printf("Error reading input: %s\nYou must enter a valid number or 0 to quit.\n", err)
			continue
		} else if bet > balance {
			fmt.Printf("Your bet of $%d exceeds your balance of $%d. Please enter a valid bet.\n", bet, balance)
			continue
		} else {
			break
		}
	}
	return bet
}
