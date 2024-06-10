package pokeapi

import "github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"

type PokemonResponse struct {
	Name            string
	Base_Experience int
	Height          int
	Weight          int
	Stats           []struct {
		Base_Stat int
		Stat      struct {
			Name string
		}
	}
}

func (pr *PokemonResponse) ToPokemon() *pokecache.Pokemon {
	stats := make([]pokecache.PokemonStat, 0, len(pr.Stats))
	for _, stat := range pr.Stats {
		stats = append(
			stats,
			pokecache.PokemonStat{
				Val:  stat.Base_Stat,
				Name: stat.Stat.Name,
			},
		)
	}
	return &pokecache.Pokemon{
		Name:   pr.Name,
		Height: pr.Height,
		Weight: pr.Weight,
		Stats:  stats,
	}
}
