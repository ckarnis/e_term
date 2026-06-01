package commands

import (
	"ecoTerm/internal/config"
	"ecoTerm/internal/windows"

	"os"

	"github.com/BurntSushi/toml"
)

func ConfigCommand(args []string) string {

	if len(args) == 0 {
		return "usage: config <edit or add>"
	}

	switch args[0] {
	case "edit":
		return edit()
	case "add":
		return add()
	case "clear":
		return clear()
	default:
		return "unknown command"
	}
}

func edit() string {
	err := windows.Manager.Open("config-add")
	if err != nil {
		return err.Error()
	}

	return ""
}

func add() string {

	return "r"
}

func clear() string {
	path := config.GetConfigPath("ecoTerm")

	cfg := config.DefaultConfig()

	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0600,
	)
	if err != nil {
		return "error"
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(cfg); err != nil {
		return "error"
	}

	return "cleared"
}
