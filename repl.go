package main

import (
	"fmt"
	"strings"

	"github.com/pablopelardas/poke-api-golang/internal/api"
)

func startRepl(config *Config) {
	scanner := *NewScanner()
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
			err := command.callback(config, words[1:])
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
	callback    func(x *Config, args []string) error
}

type Config struct {
	apiClient api.Client
	pokedex  map[string]api.Pokemon
	NextPage *string
	PrevPage *string
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
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the pokedex",
			callback:    commandPokedex,
		},
		"inspect":{
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
	}
}