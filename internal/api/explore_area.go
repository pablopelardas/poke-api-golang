package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// ExploreArea -
func (c *Client) ExploreArea(location string,) (RespLocationArea, error) {
	// Check if the pageURL is in the cache
	url := baseURL + "/location-area/" + location
	if cachedData, found := c.cache.Get(url) ; found {
		locationsResp := RespLocationArea{}
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return RespLocationArea{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)

	locationsResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	return locationsResp, nil
}