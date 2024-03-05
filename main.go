package main

import (
	"time"

	"github.com/pablopelardas/poke-api-golang/internal/api"
)

func main() {
	pokeClient := api.NewClient(5*time.Second, time.Minute*5)
	pokedex := make(map[string]api.Pokemon)
	config := &Config{
		apiClient: pokeClient,
		pokedex:   pokedex,
	}

	startRepl(config)
}
