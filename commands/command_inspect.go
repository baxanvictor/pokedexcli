package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandInspect(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	return commandWithArg("inspect", cache, displayPokemonDetails, nil, args...)
}

func displayPokemonDetails(name string, cache *pokecache.AppCache, dataHolder *any) error {
	if pokemon, found := cache.PokemonsCache.Get(name); found {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		if len(pokemon.Stats) > 0 {
			fmt.Println("Stats:")
			for _, stat := range pokemon.Stats {
				fmt.Printf("  -%v: %v\n", stat.Name, stat.Val)
			}
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
