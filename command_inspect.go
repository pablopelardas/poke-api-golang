package main

import (
	"fmt"
)

func commandInspect(config *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	if pokemon, exists := config.pokedex[args[0]]; exists {
		// print name, height, weight, stats and types
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  %s\n", t.Type.Name)
		}
		return nil
	}
	fmt.Printf("You don't have a %s\n", args[0])
	return nil


}