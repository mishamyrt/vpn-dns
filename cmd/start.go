/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"vpn-dns/internal/app"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the application in the background",
	Run: func(cmd *cobra.Command, args []string) {
		process := app.Create(configPath)
		if process.Running() {
			fmt.Println("Application is already running in background")
			return
		}
		cntxt := &daemon.Context{
			PidFileName: app.PidPath,
			PidFilePerm: 0644,
			LogFileName: app.LogPath,
			LogFilePerm: 0640,
			WorkDir:     "./",
		}

		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()

		log.Print("- - - - - - - - - - - - - - -")
		log.Print("Daemon is started")
		process.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
