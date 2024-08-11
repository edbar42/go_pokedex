package main

import (
	"sync"
	"time"

	"github.com/edbar42/go_pokedex/api"
)

// Model for application wide cache
type Cache struct {
	CacheTimeout time.Duration
	NextMap      *string
	PrevMap      *string
	MapCmdsCache map[string]MapCache
	mux          sync.Mutex
}

// Returns new application cache
func newCache(timeout time.Duration) *Cache {
	c := Cache{
		CacheTimeout: timeout,
		NextMap:      nil,
		PrevMap:      nil,
		MapCmdsCache: make(map[string]MapCache),
	}
	go c.reapLoop()
	return &c
}

// Stores the results of fetching a given URL
func (c *Cache) Add(url string, mapCache MapCache) {
	c.MapCmdsCache[url] = mapCache
}

// Retrieves a value stored in cache
func (c *Cache) Get(url string) (MapCache, bool) {
	value, ok := c.MapCmdsCache[url]
	return value, ok
}

// Clears expired entries from cache
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.CacheTimeout)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mux.Lock()
		now := time.Now()
		for k, v := range c.MapCmdsCache {
			if now.Sub(v.createdAt) > c.CacheTimeout {
				delete(c.MapCmdsCache, k)
			}
		}
		c.mux.Unlock()
	}
}

// Model for caching requests from the map and mapb commands
type MapCache struct {
	createdAt  time.Time
	cachedData api.Regions
}

// Returns map cache struct that to be stored in a map
func NewMapCache(r api.Regions) MapCache {
	return MapCache{
		createdAt:  time.Now(),
		cachedData: r,
	}
}
