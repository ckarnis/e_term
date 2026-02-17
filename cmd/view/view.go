package view

import (
	"fmt"

	"e_term/internal/query"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "view [provider] [endpoint]",
	Aliases: []string{"v"},
	Short:   "Query a configured provider",
	Args:    cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		provider := args[0]
		endpoint := args[1]

		body, err := query.QueryProvider(provider, endpoint)
		if err != nil {
			return err
		}

		fmt.Println(body)
		return nil
	},
}
