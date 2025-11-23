package main

import (
	"github.com/AsherBolleddu/pokedexcli/internal/pokeapi"
)

func commandMap(config *Config) error {
	data, err := pokeapi.GetLocationArea(config.Next)
	if err != nil {
		return err
	}

	config.Next = *data.Next
	return nil
}
