package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var ConfigPath string

func Init(appName string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appDir := filepath.Join(configDir, appName)

	if err := os.MkdirAll(appDir, 0700); err != nil {
		return err
	}

	ConfigPath = filepath.Join(appDir, "config.toml")

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		file, err := os.OpenFile(
			ConfigPath,
			os.O_CREATE|os.O_WRONLY,
			0600,
		)
		if err != nil {
			return err
		}
		defer file.Close()

		cfg := DefaultConfig()

		if err := toml.NewEncoder(file).Encode(cfg); err != nil {
			return err
		}
	}

	return nil
}
