package api

import (
	"net/http"
	"time"

	"github.com/pablopelardas/poke-api-golang/internal/cache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache 		cache.Cache
	pokedex 	map[string]Pokemon
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(cacheInterval),
	}
}