package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func startRepl() {
	commands := getCommands()
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan() // this will block the io, until it recieves a SIGTERM or user input

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if command, ok := commands[commandName]; ok {
			if err := command.Callback(); err != nil {
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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}
