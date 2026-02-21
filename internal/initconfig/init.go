package initconfig

import (
	"e_term/internal/stuff"
	"errors"
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

func InitConfigFile(force bool) error {
	if stuff.FileExists(stuff.ConfigFile) && !force {
		if !stuff.IsTerminal() {
			return errors.New("config.toml exists (use --force to overwrite)")
		}

		ok, err := confirmOverwrite()
		if err != nil {
			return err
		}
		if !ok {
			fmt.Println("")
			return nil
		}
	}
	cfg := defaultConfig()
	data, err := toml.Marshal(cfg)

	if err != nil {
		return err
	}

	return os.WriteFile(stuff.ConfigFile, data, 0644)
}
