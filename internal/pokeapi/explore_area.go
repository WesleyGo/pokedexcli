package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ExploreArea -
func (c *Client) ExploreArea(areaname *string) (LocationAreaInfo, error) {
	url := baseURL + fmt.Sprintf("/location-area/%v", *areaname)

	if c.cache != nil {
		if cached, ok := c.cache.GetFromCache(url); ok {
			locationsResp := LocationAreaInfo{}
			err := json.Unmarshal(cached, &locationsResp)
			if err != nil {
				return LocationAreaInfo{}, err
			}

			return locationsResp, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	locationsResp := LocationAreaInfo{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	if c.cache != nil {
		c.cache.AddToCache(url, dat)
	}

	return locationsResp, nil
}
