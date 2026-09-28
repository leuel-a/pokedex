package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type config struct {
	commands            map[string]cliCommand
	pokeapiClient       pokeapi.Client
	previousLocationUrl *string
	nextLocationUrl     *string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan() // this will block the io, until it recieves a SIGTERM or user input

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, exists := cfg.commands[commandName]; exists {
			if err := command.Callback(cfg); err != nil {
				fmt.Printf("[Error] error processing command: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}
