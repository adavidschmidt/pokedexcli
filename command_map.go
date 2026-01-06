package main

import (
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {

	locationPage, err := cfg.Client.FetchLocationPage(cfg.Next)

	if err != nil {
		return err
	}

	cfg.Next = locationPage.Next
	cfg.Previous = locationPage.Previous

	for _, result := range locationPage.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}

func commandMapB(cfg *Config, args ...string) error {

	if cfg.Previous == nil {
		fmt.Println("you are on the first page")
		return nil
	}

	locationPage, err := cfg.Client.FetchLocationPage(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationPage.Next
	cfg.Previous = locationPage.Previous

	for _, result := range locationPage.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}
