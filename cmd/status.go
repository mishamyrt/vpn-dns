// Package cmd contains descriptions and handlers for vpn-dns CLI.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// status represents the status command.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Prints application status",
	Run: func(cmd *cobra.Command, args []string) {
		_, daemon := createApp()
		if daemon.Running() {
			fmt.Println("Not running")
		} else {
			fmt.Println("Running")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
