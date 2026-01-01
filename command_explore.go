package main

import (
	"fmt"
)

func commandExplore(cfg *Config, location string) error {
	fmt.Printf("Exporing %s...", location)
	locationPokemon, err := cfg.Client.FetchPokemon(location)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationPokemon.EncounterList {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}
	return nil
}
