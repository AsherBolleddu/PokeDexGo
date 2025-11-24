package main

import "fmt"

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}

	locationName := args[0]
	data, err := config.pokeapiClient.GetLocationInfo(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", data.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
