package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to the Pokedex!'\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println()
	return nil
}
