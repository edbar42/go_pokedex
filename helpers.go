package main

import (
	"github.com/edbar42/go_pokedex/api"
	"github.com/edbar42/go_pokedex/ui"
)

// Helper methods for main application

// Helper for map and mapb
func mapHelper(c *api.Cache, reqURL string) error {
	cachedMap, cached := c.MapCmdsCache[reqURL]

	if cached {
		c.NextMap = cachedMap.Data.Next
		c.PrevMap = cachedMap.Data.Previous

		ui.print
	}
	// TODO: Implement this function.
	// It should:
	//		- Check for valid entries in the cache
	//		- Update the cache state if the entry is found
	//		- If it is not found, make an API request
	//		- If the api call runs or mapos are cached, print the maps to the screen
	//		- Return nil or error
	return nil
}
