package main

import (
	"fmt"
	"os"

	"github.com/edbar42/go_pokedex/api"
	"github.com/edbar42/go_pokedex/msg"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Callback    func(c *api.Cache) error
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"help": {
			Name:        "Help",
			Usage:       "help",
			Description: "Prints this help menu.",
			Callback:    help,
		},
		"exit": {
			Name:        "Exit",
			Usage:       "exit",
			Description: "Quits the program.",
			Callback:    exit,
		},
		"map": {
			Name:        "Map",
			Usage:       "map",
			Description: "Lists the next page of mapped regions.",
			Callback:    mapn,
		},
		"mapb": {
			Name:        "Map back",
			Usage:       "mapb",
			Description: "Lists the previous page of mapped regions.",
			Callback:    mapp,
		},
	}
}

func help(c *api.Cache) error {
	for _, cmd := range Commands {
		fmt.Println("------------------")
		msg.PrintCmdName(cmd.Name)
		msg.PrintCmdUsage(cmd.Usage)
		fmt.Printf("\t%s\n", cmd.Description)
		fmt.Println("------------------")
	}
	return nil
}

func exit(c *api.Cache) error {
	fmt.Println("Good luck catching them all!")
	os.Exit(0)
	return nil
}

func mapn(c *api.Cache) error {
	var reqURL string
	if c.NextMap == nil {
		reqURL = "https://pokeapi.co/api/v2/location/"
	} else {
		reqURL = *c.NextMap
	}

	// Check for entries in cache
	cachedMap, ok := c.MapCmdsCache[reqURL]
	if ok {
		c.NextMap = cachedMap.CachedData.Next
		c.PrevMap = cachedMap.CachedData.Previous

		for _, area := range cachedMap.CachedData.Results {
			msg.PrintAreaName(area.Name)
		}

		return nil
	}

	regions, err := api.FetchMappedRegions(reqURL)
	if err != nil {
		return err
	}

	c.NextMap = regions.Next
	c.PrevMap = regions.Previous

	for _, area := range regions.Results {
		msg.PrintAreaName(area.Name)
	}

	return err
}

func mapp(c *api.Cache) error {
	if c.PrevMap == nil {
		noPrevErr := msg.NoPrevMapError()
		return &noPrevErr
	}

	// Check for entries in cache
	cachedMap, ok := c.MapCmdsCache[*c.PrevMap]
	if ok {
		c.NextMap = cachedMap.CachedData.Next
		c.PrevMap = cachedMap.CachedData.Previous

		for _, area := range cachedMap.CachedData.Results {
			msg.PrintAreaName(area.Name)
		}

		return nil
	}

	regions, err := api.FetchMappedRegions(*c.PrevMap)
	if err != nil {
		return err
	}

	c.NextMap = regions.Next
	c.PrevMap = regions.Previous

	for _, area := range regions.Results {
		msg.PrintAreaName(area.Name)
	}
	return err
}
