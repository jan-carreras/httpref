package cmd

import (
	"github.com/dnnrly/httpref"
	"github.com/spf13/cobra"
)

// headersCmd represents the headers command
var headersCmd = &cobra.Command{
	Use:     "headers",
	Aliases: []string{"header"},
	Short:   "References for common HTTP headers",
	Run: func(cmd *cobra.Command, args []string) {
		results := httpref.Headers

		if len(args) == 0 {
			results = results.Titles()
		} else {
			results = results.ByName(args[0])
		}

		printResults(results)
	},
}

func init() {
	rootCmd.AddCommand(headersCmd)
}
