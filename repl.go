package main

import (
	"fmt"
	"strings"
)

func startRepl() {
	scanner := *NewScanner()
	config := &Config{
		NextPage: "",
		PrevPage: "",
	}
	defer scanner.Close()

	for {
		scanner.Prompt("pokedex> ")
		input := scanner.Scan()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(x *Config) error
}

type Config struct {
	NextPage string
	PrevPage string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of the map",
			callback:    commandMapB,
		},
	}
}