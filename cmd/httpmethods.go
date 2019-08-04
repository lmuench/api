package cmd

import (
	"fmt"

	"github.com/lmuench/api/api"
	"github.com/spf13/cobra"
)

func init() {
	registerHttpMethods()
}

func registerHttpMethods() {
	methods := []string{"POST", "GET", "PUT", "PATCH", "DELETE"}
	for _, method := range methods {
		rootCmd.AddCommand(
			&cobra.Command{
				Use:   fmt.Sprintf("%s [URL] [body]", method),
				Short: fmt.Sprintf("HTTP %s request", method),
				Long:  fmt.Sprintf("Execute an HTTP %s request against the given URL", method),
				Run: func(cmd *cobra.Command, args []string) {
					call := api.Call{
						Command: method,
						Args:    args,
					}
					answer := api.Handle(call)
					fmt.Println(answer.Result)
				},
			},
		)
	}
}
