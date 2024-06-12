package pokeapi

type Config struct {
	Previous string
	Next     string
}

type LocationAreasResponse struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func (c *Config) UpdateFromResponse(response *LocationAreasResponse) {
	c.Next = response.Next
	c.Previous = response.Previous
}
