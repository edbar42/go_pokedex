package ui

import (
	"fmt"

	"github.com/edbar42/go_pokedex/api"
)

// Custom printings for the help menu
func PrintCmdUsage(s string) {
	fmt.Printf("\t\x1b[32m%s\x1b[0m\n", s)
}

func PrintCmdName(s string) {
	fmt.Printf("\x1b[36m%s\x1b[0m\n", s)
}

// Custom printing for area names in mapped regions
func PrintAreaName(s string) {
	fmt.Printf("\x1b[1;35m%s\x1b[0m\n", s)
}

// Helper function for map commands outputs
func PrintRegions(c api.MapCache) {
	for _, area := range c.Data.Results {
		PrintAreaName(area.Name)
	}
}
