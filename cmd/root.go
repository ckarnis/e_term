package cmd

import (
	"e_term/cmd/view"
	"e_term/cmd/workspace"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eterm",
	Short: "eterm is a multi-tool CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(view.Cmd)
	rootCmd.AddCommand(workspace.InitCmd)
	rootCmd.AddCommand(workspace.AddCmd)
	rootCmd.AddCommand(workspace.EditCmd)
}
