package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// runCmd represents the run command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the application in the current process",
	Run: func(cmd *cobra.Command, args []string) {
		app, daemon := createApp()
		if daemon.Running() {
			fmt.Println("Application is running in background. Stop it, before run")
			os.Exit(1)
		}
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
