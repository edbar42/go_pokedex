// Custom terminal output for the application
package ui

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
