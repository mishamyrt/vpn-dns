package cmd

import (
	"fmt"
	"vpn-dns/internal/app"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the background application",
	Run: func(cmd *cobra.Command, args []string) {
		process := app.Create(configPath)
		if process.Running() {
			process.Kill()
			fmt.Println("Daemon is stopped")
		} else {
			fmt.Println("Daemon is not running")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
