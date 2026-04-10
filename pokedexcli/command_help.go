package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	reset   = "\033[0m"
	bold    = "\033[1m"
	dim     = "\033[2m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
)

func useColor() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	term := os.Getenv("TERM")
	return term != "" && term != "dumb"
}

func c(style, s string) string {
	if useColor() {
		return style + s + reset
	}
	return s
}

const pikachu = `
               (\_/)
               (•_•)   P O K É D E X
            __/ >🍩   "gotta fetch 'em all"
`
const version = "1.0.0"

func renderHelp() {
	title := c(bold+yellow, "pokedexcli")
	bar := c(blue, strings.Repeat("─", 48))

	fmt.Println()
	fmt.Printf("%s %s  %s%s%s\n", c(bold+magenta, "⚡"), title, c(dim, "v"), c(dim, version), c(dim, " — CLI Companion"))
	fmt.Println(bar)
	fmt.Print(c(yellow, pikachu))
	fmt.Println(bar)

	fmt.Printf("%s\n", c(bold, "Usage"))
	fmt.Printf("  %s %s\n\n", c(green, "./pokedexcli\n"), c(cyan, "[command] [flags]"))
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("  %s  %s\n", c(green, cmd.name), c(dim, cmd.description))
	}
	fmt.Println()

	fmt.Printf("%s\n", c(bold, "Tips"))
	fmt.Printf("  %s  type %s to see this panel\n", c(dim, "• In POKEDEX,"), c(green, "help"))
	fmt.Printf("  %s  press %s to exit\n", c(dim, "• Anytime,"), c(green, "Ctrl+C"))
	fmt.Println(bar)
	fmt.Println()
}

func callbackHelp(cfg *config) error {
	renderHelp()
	return nil
}
