package main

import (
	"errors"
	"fmt"
	"os"
	"poke_api"
)

func commandHelp(config *poke_api.Config) error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, command := range cliCommands() {
		fmt.Printf("%v: %v\n", name, command.description)
	}

	fmt.Println()

	return nil
}

func commandExit(config *poke_api.Config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *poke_api.Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != nil {
		url = *config.Next
	}
	return requestLocationAreas(url, config)
}

func commandMapB(config *poke_api.Config) error {
	if config.Previous == nil {
		return errors.New("don't have what to go back to")
	}
	return requestLocationAreas(*config.Previous, config)
}

func requestLocationAreas(url string, config *poke_api.Config) error {
	fmt.Printf("Requesting %s...\n", url)
	response := poke_api.LocationAreasResponse{}
	err := poke_api.SendGetRequest(url, &response)
	if err != nil {
		return err
	}
	config.Next = response.Next
	config.Previous = response.Previous
	fmt.Println("Received location areas:")
	for _, area := range response.Results {
		if area.Name != nil {
			fmt.Println(*area.Name)
		}
	}
	fmt.Println()
	return nil
}
