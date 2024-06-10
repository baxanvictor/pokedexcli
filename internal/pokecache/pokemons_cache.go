package pokecache

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  []PokemonStat
}

type PokemonStat struct {
	Val  int
	Name string
}
