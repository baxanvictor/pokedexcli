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

	for name, command := range CliCommands() {
		fmt.Printf("%v: %v\n", name, command.description)
	}

	fmt.Println()

	return nil
}
