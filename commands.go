package main

import (
	"errors"
	"fmt"
	"internal/pokeapi"
	"internal/pokecache"
	"os"
)

func commandHelp(config *pokeapi.Config, cache *pokecache.Cache) error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, command := range cliCommands() {
		fmt.Printf("%v: %v\n", name, command.description)
	}

	fmt.Println()

	return nil
}

func commandExit(config *pokeapi.Config, cache *pokecache.Cache) error {
	os.Exit(0)
	return nil
}

func commandMap(config *pokeapi.Config, cache *pokecache.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != nil {
		url = *config.Next
	}
	return requestLocationAreas(url, config, cache)
}

func commandMapB(config *pokeapi.Config, cache *pokecache.Cache) error {
	if config.Previous == nil {
		return errors.New("don't have what to go back to")
	}
	return requestLocationAreas(*config.Previous, config, cache)
}

func requestLocationAreas(url string, config *pokeapi.Config, cache *pokecache.Cache) error {
	response := pokeapi.LocationAreasResponse{}

	if cacheHit := tryHitCache(url, cache, &response); !cacheHit {
		fmt.Printf("Requesting %s...\n", url)
		body, err := pokeapi.SendGetRequest(url, &response)
		if err != nil {
			return err
		}
		cache.Add(url, body)
	}

	updateConfigAndPrintResults(config, &response)
	return nil
}

func tryHitCache(
	url string,
	cache *pokecache.Cache,
	response *pokeapi.LocationAreasResponse,
) (cacheHit bool) {
	if currentData, ok := cache.Get(url); ok {
		fmt.Printf("Found data in cache %s...\n", url)
		err := pokeapi.UnmarshalResponseBody(currentData, &response)
		if err != nil {
			cache.Remove(url)
		} else {
			cacheHit = ok
		}
	}
	return cacheHit
}

func updateConfigAndPrintResults(config *pokeapi.Config, response *pokeapi.LocationAreasResponse) {
	config.UpdateFromResponse(response)
	fmt.Println("Received location areas:")
	for _, area := range response.Results {
		if area.Name != nil {
			fmt.Println(*area.Name)
		}
	}
	fmt.Println()
}
