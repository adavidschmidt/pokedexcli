package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(location string) (LocationPokemon, error) {
	url := baseURL + "/location-area/" + location
	var data []byte

	entry, ok := c.cache.Get(url)
	if !ok {

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return LocationPokemon{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationPokemon{}, err
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationPokemon{}, err
		}
		c.cache.Add(url, data)
	} else {
		data = entry
	}

	var locationPokemon LocationPokemon

	if err := json.Unmarshal(data, &locationPokemon); err != nil {
		return LocationPokemon{}, fmt.Errorf("error decoding response body: %w", err)
	}
	return locationPokemon, nil
}
