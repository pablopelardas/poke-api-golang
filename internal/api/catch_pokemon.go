package api

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) CatchPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	req, err := http.NewRequest("GET", url, nil)
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	catchResp := Pokemon{}
	err = json.Unmarshal(dat, &catchResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return catchResp, nil
}