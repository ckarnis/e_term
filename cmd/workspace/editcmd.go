package workspace

import (
	"e_term/internal/editconfig"

	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a data source",
	Run: func(cmd *cobra.Command, args []string) {
		editconfig.EditSource()
	},
}
