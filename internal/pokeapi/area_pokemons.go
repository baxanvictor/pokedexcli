package pokeapi

type AreaPokemonsResponse struct {
	Pokemon_Encounters []struct {
		Pokemon struct {
			Name *string
			Url  *string
		}
	}
}
