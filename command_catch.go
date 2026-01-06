package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemon, err := cfg.Client.FetchPokemonInfo(args[0])
	if err != nil {
		return err
	}
	baseExp := pokemon.BaseExperience
	randomNum := rand.Intn(baseExp)
	if randomNum > 40 {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.CaughtPokemon[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command")
		return nil
	} else {
		fmt.Printf("%s escaped!\n", args[0])
		return nil
	}
}
