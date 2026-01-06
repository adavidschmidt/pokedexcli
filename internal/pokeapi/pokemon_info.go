package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemonInfo(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	var data []byte

	entry, ok := c.cache.Get(url)
	if !ok {

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return Pokemon{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, data)
	} else {
		data = entry
	}

	var catchingPokemon Pokemon

	if err := json.Unmarshal(data, &catchingPokemon); err != nil {
		return Pokemon{}, fmt.Errorf("error decoding response body: %w", err)
	}
	return catchingPokemon, nil
}
