// config/path.go
package config

import (
	"os"
	"path/filepath"
)

func GetConfigPath(appName string) string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, appName, "config.toml")
}
