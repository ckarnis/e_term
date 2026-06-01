package commands

import (
	"ecoTerm/internal/windows"
)

func Open(args []string) string {

	if len(args) == 0 {
		return "usage: open <name>"
	}

	err := windows.Manager.Open(args[0])

	if err != nil {
		return err.Error()
	}

	return "opened " + args[0]
}

func Close(args []string) string {

	if len(args) == 0 {
		return "usage: close <name>"
	}

	err := windows.Manager.Close(args[0])

	if err != nil {
		return err.Error()
	}

	return "closed " + args[0]
}
