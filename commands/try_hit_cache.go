package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func tryHitCache[T any](
	url string,
	cache *pokecache.Cache,
	response *T,
) (cacheHit bool) {
	if currentData, ok := cache.Get(url); ok {
		fmt.Printf("Found data in cache %s...\n", url)
		err := pokeapi.UnmarshalResponseBody(currentData, &response)
		if err != nil {
			cache.Remove(url)
		} else {
			cacheHit = ok
		}
	}
	return cacheHit
}
