package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationAreaDetail(LocationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/"
	fullURL := baseUrl + endpoint + LocationAreaName

 
	// check cache
	dat,ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("found cache")
		locationAreaDetailResp := LocationArea{}
		err := json.Unmarshal(dat, &locationAreaDetailResp)
		if err != nil {
		return LocationArea{}, err
		}
		return locationAreaDetailResp, nil
	}
	//requsting
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationArea{}, err
	}
	//decoding response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("error with status of %v", resp.StatusCode)
	}

	// reading response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	// cacheing data
	c.cache.Add(fullURL,data)
	fmt.Println("new cache")
	// set data with generated struct, using Unmarshal
	locationAreaDetailResp := LocationArea{}
	err = json.Unmarshal(data, &locationAreaDetailResp)
	if err != nil {
		return LocationArea{}, err
	}
	return locationAreaDetailResp, nil
}

