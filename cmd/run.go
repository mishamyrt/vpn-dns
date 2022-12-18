package cmd

import (
	"fmt"

	"vpn-dns/internal/app"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the application in the current process",
	Run: func(cmd *cobra.Command, args []string) {
		process := app.Create(configPath)
		if process.Running() {
			fmt.Println("Application is running in background. Stop it, before run")
			return
		}
		process.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
