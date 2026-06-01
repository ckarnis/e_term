package commands

// where the inputs
// are read into commands

import (
	"strings"
)

func Execute(input string) string {

	parts := strings.Fields(input)

	if len(parts) == 0 {
		return ""
	}

	switch parts[0] {

	case "open":
		return Open(parts[1:])

	case "close":
		return Close(parts[1:])

	case "help":
		return Help()

	case "config":
		return ConfigCommand(parts[1:])

	default:
		return "unknown command"
	}
}
