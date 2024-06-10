package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func tryHitCache[T any](
	url string,
	cache *pokecache.AppCache,
	response *T,
) (cacheHit bool) {
	if currentData, ok := cache.Cache.Get(url); ok {
		fmt.Printf("Found data in cache %s...\n", url)
		err := pokeapi.UnmarshalResponseBody(currentData, &response)
		if err != nil {
			cache.Cache.Remove(url)
		} else {
			cacheHit = ok
		}
	}
	return cacheHit
}
