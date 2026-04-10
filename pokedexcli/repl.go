package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const pokeBall = "◓"

func gitBranch() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	headPath := filepath.Join(wd, ".git", "HEAD")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return ""
	}
	s := strings.TrimSpace(string(data))
	const prefix = "ref: refs/heads/"
	if strings.HasPrefix(s, prefix) {
		return strings.TrimPrefix(s, prefix)
	}
	if len(s) >= 7 {
		return s[:7]
	}
	return s
}

func cwdBase() string {
	wd, err := os.Getwd()
	if err != nil {
		return "?"
	}
	return filepath.Base(wd)
}

func renderPrompt() string {
	left := c(bold+magenta, "["+pokeBall+" ") +
		c(bold+yellow, "pokedexcli ") +
		c(dim, "v"+version) +
		c(bold+magenta, "]")

	dir := c(bold+blue, "["+cwdBase()+"]")

	br := gitBranch()
	branch := ""
	if br != "" {
		branch = " " + c(cyan, "("+br+")")
	}

	t := c(dim, time.Now().Format("15:04"))

	arrow := c(green, " ➤ ")

	return left + " " + dir + branch + "  " + t + arrow
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists some Poke Locations Areas",
			callback:    CallbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous Poke Locations Areas",
			callback:    CallbackMapb,
		},

		"exit": {
			name:        "exit",
			description: "Exits the Pokedex CLI",
			callback:    callbackExit,
		},
	}
}

func startRepl(cfg *config) {
	renderHelp()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(renderPrompt())
		scanner.Scan()
		text := scanner.Text()
		cleaned := CleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		cmdName := cleaned[0]
		availableCommands := getCommands()

		cmd, ok := availableCommands[cmdName]
		if !ok {
			fmt.Printf("Unrecognized command: %s\n", cmdName)
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Printf("Error executing command %s: %v\n", cmdName, err)
		}
	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
