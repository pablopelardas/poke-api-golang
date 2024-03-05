package main

import (
	"fmt"
	"math/rand"
)


func commandCatch(config *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}
	// check if pokemon is in pokedex
	if _, exists := config.pokedex[args[0]]; exists {
		fmt.Printf("You already have a %s\n", args[0])
		return nil
	}
	fmt.Printf("Catching pokemon... %s", args[0])
	fmt.Println()

	pokemon, err := config.apiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}
	// Randomly decide if the pokemon was caught with probability depending on the pokemon's base experience
	random := rand.Intn( pokemon.BaseExperience )
	if random < pokemon.BaseExperience*3/4 {
			fmt.Printf("The %s got away!\n", pokemon.Name)
			return nil	
	}
	fmt.Printf("You caught a %s!\n", pokemon.Name)

	config.pokedex[pokemon.Name] = pokemon

	return nil
}