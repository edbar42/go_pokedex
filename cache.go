package main

import (
	"time"

	"github.com/edbar42/go_pokedex/api"
)

// Model for application wide cache
type Cache struct {
	NextMap      *string
	PrevMap      *string
	MapCmdsCache map[string]MapCache
}

// Model for caching requests from the map and mapb commands
type MapCache struct {
	createdAt  time.Time
	cachedData api.Regions
}

// TODO: Figure out how time.Time and time.Duration relate to each other
func newMapCache(timeout time.Duration, r api.Regions) MapCache {
	return MapCache{
		createdAt: timeout,
	}
}

func newCache() Cache {
	return Cache{}
}
