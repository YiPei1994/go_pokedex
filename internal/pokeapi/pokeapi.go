package pokeapi

import (
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

// creating client struct and create new client object

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
