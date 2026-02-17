package workspace

import (
	"e_term/internal/config"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new data source",
	Run: func(cmd *cobra.Command, args []string) {
		config.AddSource()
	},
}
