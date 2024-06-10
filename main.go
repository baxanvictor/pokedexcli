package main

import (
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func main() {
	appCache := pokecache.AppCache{
		Cache:         pokecache.NewCache[[]byte](nil, true),
		PokemonsCache: pokecache.NewCache[pokecache.Pokemon](nil, false),
	}
	startRepl(&pokeapi.Config{}, &appCache)
}
