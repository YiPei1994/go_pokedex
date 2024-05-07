package pokeapi

import (
	"net/http"
	"time"

	"github.com/YiPei1994/PokeGo/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

// creating client struct and create new client object

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
