package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	data, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil

}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("you are on the first page")
	}
	data, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}