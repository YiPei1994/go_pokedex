package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationsAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area/"
	fullURL := baseUrl + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	//requsting
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}
	//decoding response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("error with status of %v", resp.StatusCode)
	}

	// reading response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// set data with generated struct, using Unmarshal
	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	return locationAreaResp, nil
}
