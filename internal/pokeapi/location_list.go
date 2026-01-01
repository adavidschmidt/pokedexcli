package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocationPage(pageURL *string) (LocationPage, error) {
	url := baseURL + "/location-area"
	var data []byte
	if pageURL != nil {
		url = *pageURL
	}
	entry, ok := c.cache.Get(url)
	if !ok {

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return LocationPage{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationPage{}, err
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationPage{}, err
		}
		c.cache.Add(url, data)
	} else {
		data = entry
	}

	var locationPage LocationPage

	if err := json.Unmarshal(data, &locationPage); err != nil {
		return LocationPage{}, fmt.Errorf("error decoding response body: %w", err)
	}
	return locationPage, nil
}
