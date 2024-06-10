package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandExplore(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	return commandWithArg(
		"explore",
		cache,
		requestPokemonsInArea,
		&pokeapi.AreaPokemonsResponse{},
		args...,
	)
}

func requestPokemonsInArea(areaName string, cache *pokecache.AppCache, response *pokeapi.AreaPokemonsResponse) error {
	fmt.Printf("Exploring %s...\n", areaName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)

	if cacheHit := tryHitCache(url, cache, &response); !cacheHit {
		fmt.Printf("Requesting %s...\n", url)
		body, err := pokeapi.SendGetRequest(url, &response)
		if err != nil {
			return err
		}
		cache.Cache.Add(url, &body)
	}

	fmt.Println("Found Pokemons:")
	for _, encounter := range response.Pokemon_Encounters {
		fmt.Println(*encounter.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
