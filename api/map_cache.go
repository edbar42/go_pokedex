package api

import (
	"time"
)

// Model for caching requests from the map and mapb commands
type MapCache struct {
	CreatedAt time.Time
	Data      Regions
}

// Returns cached regions names from mapping commands
func NewMapCache(r Regions) MapCache {
	return MapCache{
		CreatedAt: time.Now(),
		Data:      r,
	}
}

// Stores the results of fetching a given URL
func (c *Cache) AddMap(url string, mapCache MapCache) {
	c.MapCmdsCache[url] = mapCache
}

// Retrieves a value stored in cache
func (c *Cache) GetMap(url string) (MapCache, bool) {
	value, ok := c.MapCmdsCache[url]
	return value, ok
}
