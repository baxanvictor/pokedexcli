package commands

import (
	"fmt"
	"math/rand"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandCatch(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	response := pokeapi.PokemonResponse{}
	if err := commandWithArg(
		"catch",
		cache,
		requestPokemonDetails,
		&response,
		args...,
	); err != nil {
		return err
	}

	if ball := rand.Intn(255); ball >= *response.Base_Experience {
		cache.PokemonsCache.Add(response.Name, &pokecache.Pokemon{Name: response.Name})
		fmt.Printf("%s was caught!\n", response.Name)
	} else {
		fmt.Printf("%s escaped!\n", response.Name)
	}
	return nil
}

func requestPokemonDetails(pokemonName string, cache *pokecache.AppCache, response *pokeapi.PokemonResponse) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

	if cacheHit := tryHitCache(url, cache, &response); !cacheHit {
		fmt.Printf("Requesting %s...\n", url)
		body, err := pokeapi.SendGetRequest(url, &response)
		if err != nil {
			return err
		}
		cache.Cache.Add(url, &body)
	}

	return nil
}
