package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name\n")
	}

	pokemonName := args[0]

	if _, ok := config.caughtPokemon[pokemonName]; ok {
		fmt.Printf("%s already caught", pokemonName)
		return nil
	}

	pokemon, err := config.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
