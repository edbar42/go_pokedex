// Custom terminal output for the application
package msg

import "fmt"

const Prompt = "\033[5m|> \033[0m"

// Custom errors for the application
type PokeError struct {
	Msg string
}

func (pe *PokeError) Error() string {
	return fmt.Sprintf("\033[31m%s\033[0m", pe.Msg)
}

// Custom error type for command not found
func CmdNotFoundErr() PokeError {
	return PokeError{
		Msg: "Hmm... I am not sure what you mean. Try `help` if you are feeling lost.",
	}
}

// Custom error for mapping back while at the first list of regions
func NoPrevMapError() PokeError {
	return PokeError{
		Msg: "There are no regions to list back to.",
	}
}

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
