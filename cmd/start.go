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
		app, daemon := createApp()
		if daemon.Running() {
			fmt.Println("Application is already running in background")
			os.Exit(1)
		}
		child, err := daemon.Context.Reborn()
		if err != nil {
			fmt.Println("Error while starting daemon:", err.Error())
		}
		if child == nil {
			defer daemon.Context.Release()
			app.Run()
		} else {
			fmt.Println("Daemon is started")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
