package main

import (
	"fmt"
)

func commandPokedex(config *Config, _ []string) error {
	fmt.Println("Your pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}