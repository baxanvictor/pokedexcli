package pokecache

type AppCache struct {
	Cache         *Cache[[]byte]
	PokemonsCache *Cache[Pokemon]
}

func (ac *AppCache) ListPokemonNames() []string {
	names := make([]string, 0, len(ac.PokemonsCache.Entries))
	for _, cacheEntry := range ac.PokemonsCache.Entries {
		names = append(names, cacheEntry.val.Name)
	}
	return names
}
