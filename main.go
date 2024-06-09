package main

import (
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(nil)
	startRepl(&pokeapi.Config{}, cache)
}
