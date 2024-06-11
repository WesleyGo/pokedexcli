package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationInfo, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if c.cache != nil {
		if cached, ok := c.cache.GetFromCache(url); ok {
			locationsResp := LocationInfo{}
			err := json.Unmarshal(cached, &locationsResp)
			if err != nil {
				return LocationInfo{}, err
			}

			return locationsResp, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, err
	}

	locationsResp := LocationInfo{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationInfo{}, err
	}

	if c.cache != nil {
		c.cache.AddToCache(url, dat)
	}

	return locationsResp, nil
}
