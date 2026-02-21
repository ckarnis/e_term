package workspace

import (
	"e_term/internal/initconfig"

	"github.com/spf13/cobra"
)

var force bool

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a config.toml file in the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		return initconfig.InitConfigFile(force)
	},
}

func init() {
	InitCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing config.toml")
}
