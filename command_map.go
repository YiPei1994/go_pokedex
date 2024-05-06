package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)
type LocationAreaResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func callMapLocations() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
	  	return errors.New("something went wrong")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var data LocationAreaResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return errors.New("something went wrong")
	}
	for _, location := range data.Results {
		fmt.Println( location.Name)
	}

	return nil
}