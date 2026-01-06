package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	Client        pokeapi.Client
	Next          *string
	Previous      *string
	CaughtPokemon map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Provide a list of commands and their description",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Provide a list of location areas. Will move to the next page of location areas if used consecutively",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Provides the previous page of location area results, or indicates you are on the first page",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Provides a list of pokemon that can be encountered in a give location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Makes an attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Provides information on the given pokemon if the user has caught the given pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Provides the user with a list of their caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Exiting....")
			break
		}
		text := scanner.Text()
		cleanText := cleanInput(text)
		commandName := cleanText[0]
		args := cleanText[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	textList := strings.Fields(lowered)
	return textList
}
