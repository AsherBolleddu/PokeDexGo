package main

import "fmt"

func commandInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide pokemon name\n")
	}

	pokemonName := args[0]

	pokemon, ok := config.caughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, value := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range pokemon.Types {
		fmt.Printf("  - %s\n", value.Type.Name)
	}

	return nil
}
