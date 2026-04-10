package main

import (
	"fmt"
	"os"
	"strings"
)

func callbackExit(cfg *config) error {
	title := c(bold+yellow, "pokedexcli")
	bar := c(blue, strings.Repeat("─", 48))

	sassyPika := c(cyan, `
                 (\_/)
                 (•_•)   P I K A  ...
               __/ >🥺   "leaving already?"
    `)

	fmt.Println()
	fmt.Println(bar)
	fmt.Printf("%s %s\n", c(bold+magenta, "⚡"), title)
	fmt.Println(bar)

	fmt.Println(sassyPika)

	fmt.Println(c(yellow, "    Hmph. Fine. Go."))
	fmt.Println(c(dim, "    • Just know *I'm* not the one missing out..."))
	fmt.Println()

	fmt.Println(bar)
	fmt.Println(c(bold+green, "Pokedex CLI: Logged out — but I’ll be stronger when you return."))
	fmt.Println(bar)

	os.Exit(0)
	return nil
}
