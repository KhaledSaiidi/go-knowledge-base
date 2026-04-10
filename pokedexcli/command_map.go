package main

import (
	"fmt"
	"strings"
)

const (
	red = "\033[31m"
)

func CallbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	title := c(bold+yellow, "Available Location Areas")
	bar := c(blue, strings.Repeat("─", 48))
	count := c(dim, fmt.Sprintf("Total: %d", resp.Count))

	fmt.Println()
	fmt.Println(title)
	fmt.Println(bar)
	fmt.Println(count)
	fmt.Println()

	for i, area := range resp.Results {
		line := fmt.Sprintf("%2d. %s", i+1, c(green, area.Name))
		fmt.Println("  " + line)
	}

	fmt.Println()
	fmt.Println(bar)
	fmt.Println()
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func CallbackMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		fmt.Println()
		fmt.Println(c(red, "No previous page available"))
		fmt.Println()
		return nil
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	title := c(bold+yellow, "Available Location Areas")
	bar := c(blue, strings.Repeat("─", 48))
	count := c(dim, fmt.Sprintf("Total: %d", resp.Count))

	fmt.Println()
	fmt.Println(title)
	fmt.Println(bar)
	fmt.Println(count)
	fmt.Println()

	for i, area := range resp.Results {
		line := fmt.Sprintf("%2d. %s", i+1, c(green, area.Name))
		fmt.Println("  " + line)
	}

	fmt.Println()
	fmt.Println(bar)
	fmt.Println()
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
