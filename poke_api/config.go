package poke_api

type Config struct {
	Previous *string
	Next     *string
}

type LocationAreasResponse struct {
	Next     *string
	Previous *string
	Results  []struct {
		Name *string
		Url  *string
	}
}
