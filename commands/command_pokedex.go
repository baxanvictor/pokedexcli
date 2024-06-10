package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandPokedex(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, name := range cache.ListPokemonNames() {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
