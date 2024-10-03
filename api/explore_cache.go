package api

import (
	"time"
)

// Model for caching requests from the map and mapb commands
type ExploreCache struct {
	CreatedAt time.Time
	Data      Area
}

// Returns cached regions names from mapping commands
func NewExploreCache(a Area) ExploreCache {
	return ExploreCache{
		CreatedAt: time.Now(),
		Data:      a,
	}
}

// Stores the results of fetching a given URL
func (c *Cache) AddExplore(areaName string, exploreCache ExploreCache) {
	c.ExploreCmdsCache[areaName] = exploreCache
}

// Retrieves a value stored in cache
func (c *Cache) GetExplore(areaName string) (ExploreCache, bool) {
	value, ok := c.ExploreCmdsCache[areaName]
	return value, ok
}
