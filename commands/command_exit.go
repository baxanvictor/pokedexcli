package commands

import (
	"os"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandExit(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	os.Exit(0)
	return nil
}
