package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		words := cleanInput(scanner.Text())

		if len(words) != 0 {
			fmt.Printf("Your command was: %s\n", words[0])
		}

		fmt.Print("Pokedex > ")
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)

	return strings.Fields(lower)
}
