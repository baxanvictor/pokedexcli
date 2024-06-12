package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandHelp(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for pair := CliCommands().Oldest(); pair != nil; pair = pair.Next() {
		fmt.Printf("%v: %v\n", pair.Key, pair.Value.(*cliCommand).description)
	}

	fmt.Println()

	return nil
}
