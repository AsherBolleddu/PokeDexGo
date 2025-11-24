package main

import "fmt"

func commandPokedex(config *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
