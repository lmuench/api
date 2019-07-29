package cmd

import (
	"fmt"

	"github.com/lmuench/api/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "GET [URL]",
	Short: "HTTP GET request",
	Long:  `Execute an HTTP GET request against the given URL`,
	Run: func(cmd *cobra.Command, args []string) {
		call := api.Call{
			Command: "GET",
			Args:    args,
		}
		answer := api.Handle(call)
		fmt.Println(answer.Result)
	},
}
