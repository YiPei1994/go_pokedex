package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (PokemonDetail, error) {
	endpoint := "/pokemon/"
	fullURL := baseUrl + endpoint + pokemonName

 
	// check cache
	dat,ok := c.cache.Get(fullURL)
	if ok {
		
		pokemonDetail := PokemonDetail{}
		err := json.Unmarshal(dat, &pokemonDetail)
		if err != nil {
		return PokemonDetail{}, err
		}
		return pokemonDetail, nil
	}
	//requsting
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return PokemonDetail{}, err
	}
	//decoding response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetail{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return PokemonDetail{}, fmt.Errorf("error with status of %v", resp.StatusCode)
	}

	// reading response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetail{}, err
	}
	// cacheing data
	c.cache.Add(fullURL,data)
	
	// set data with generated struct, using Unmarshal
	pokemonDetail := PokemonDetail{}
	err = json.Unmarshal(data, &pokemonDetail)
	if err != nil {
		return PokemonDetail{}, err
	}
	return pokemonDetail, nil
}

