package ui

import "fmt"

// Custom printing for the help menu
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
