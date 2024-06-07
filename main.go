package main

import (
	"internal/pokeapi"
	"internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(nil)
	startRepl(&pokeapi.Config{}, cache)
}
