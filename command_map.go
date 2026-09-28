package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationRes.Next
	cfg.previousLocationUrl = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationRes.Next
	cfg.previousLocationUrl = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
