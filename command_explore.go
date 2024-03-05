package main

import (
	"fmt"
)

func commandExplore(config *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location>")
	}
	locationsResp, err := config.apiClient.ExploreArea(args[0])
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Exploring %s, Found Pokemon: \n", args[0])
	for _, pokemon := range locationsResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}