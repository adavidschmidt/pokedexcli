package main

import (
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("No arguments provided. Please provide a location to explore")
		return nil
	}
	fmt.Printf("Exploring %s...", args[0])
	locationPokemon, err := cfg.Client.FetchPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationPokemon.EncounterList {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
