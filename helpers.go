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

		ui.PrintRegions(cachedMap)
		return nil
	}

	regions, err := api.FetchMappedRegions(reqURL)
	if err != nil {
		return err
	}

	fetched := api.NewMapCache(regions)

	c.Add(reqURL, fetched)
	c.NextMap = regions.Next
	c.PrevMap = regions.Previous

	ui.PrintRegions(fetched)

	return nil
}
