package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the application in the background",
	Run: func(cmd *cobra.Command, args []string) {
		changer, daemon := createApp()
		if daemon.Running() {
			fmt.Println("Application is already running in background")
			os.Exit(1)
		}
		daemon.Start()
		changer.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
