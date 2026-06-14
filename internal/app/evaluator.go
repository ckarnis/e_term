package app

import (
	"ecoTerm/internal/commands"
	"strings"
)

func execute(input string) string {

	parts := strings.Fields(input)

	if len(parts) == 0 {
		return ""
	}

	switch parts[0] {

	case "open":
		return commands.Open(parts[1:])

	case "close":
		return commands.Close(parts[1:])

	case "help":
		return commands.Help()

	case "config":
		return commands.ConfigCommand(parts[1:])

	case "view":
		return commands.View(parts[1:])

	default:
		return "unknown command"
	}
}
