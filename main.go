package main

import (
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func main() {
	appCache := pokecache.AppCache{
		Cache:         pokecache.NewCache[[]byte](0, true),
		PokemonsCache: pokecache.NewCache[pokecache.Pokemon](0, false),
	}
	startRepl(&pokeapi.Config{}, &appCache)
}
