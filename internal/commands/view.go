package commands

import (
	"ecoTerm/internal/config"
	"ecoTerm/internal/view"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func View(args []string) string {
	if len(args) == 0 {
		return "usage: view <source> <type> <dataset>"
	}

	var fileData map[string]any

	path := config.GetConfigPath("ecoTerm")

	if _, err := toml.DecodeFile(path, &fileData); err != nil {
		fmt.Println("can't read config.toml")
		os.Exit(1)
	}

	item, ok := fileData[args[0]]
	if !ok {
		return fmt.Sprintf("source %q not found", args[0])
	}

	switch args[0] {
	case "fred":
		return view.ViewFred(args[1:], item)
	default:
		return "unknow source"
	}
}
