package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ExploreArea -
func (c *Client) CatchPokemon(pokemon *string) (PokemonStats, error) {
	url := baseURL + fmt.Sprintf("/pokemon/%v", *pokemon)

	if c.cache != nil {
		if cached, ok := c.cache.GetFromCache(url); ok {
			stats := PokemonStats{}
			err := json.Unmarshal(cached, &stats)
			if err != nil {
				return PokemonStats{}, err
			}

			return stats, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonStats{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonStats{}, err
	}

	stats := PokemonStats{}
	err = json.Unmarshal(dat, &stats)
	if err != nil {
		return PokemonStats{}, err
	}

	if c.cache != nil {
		c.cache.AddToCache(url, dat)
	}

	return stats, nil
}
