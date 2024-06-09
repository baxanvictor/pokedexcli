package commands

import (
	"errors"
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandExplore(config *pokeapi.Config, cache *pokecache.Cache, args ...string) error {
	switch argsLen := len(args); argsLen {
	case 0:
		return errors.New("an area name must be provided with the \"explore\" command")
	case 1:
		return requestPokemonsInArea(args[0], cache)
	default:
		return errors.New("the \"explore\" command only takes one argument")
	}
}

func requestPokemonsInArea(areaName string, cache *pokecache.Cache) error {
	fmt.Printf("Exploring %s...\n", areaName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)
	response := pokeapi.AreaPokemonsResponse{}

	if cacheHit := tryHitCache(url, cache, &response); !cacheHit {
		fmt.Printf("Requesting %s...\n", url)
		body, err := pokeapi.SendGetRequest(url, &response)
		if err != nil {
			return err
		}
		cache.Add(url, body)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range response.Pokemon_Encounters {
		fmt.Println(*encounter.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
