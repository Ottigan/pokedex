package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var locationUrl = "https://pokeapi.co/api/v2/location-area"

func main() {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)

	config := &config{
		Next:     &locationUrl,
		Previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		cmd, ok := commands[input[0]]

		if ok {
			if err := cmd.callback(config); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

var commands = map[string]cliCommand{}

func initCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)

	return strings.Fields(lower)
}

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
