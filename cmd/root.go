// Package cmd contains descriptions and handlers for vpn-dns CLI.
package cmd

import (
	"fmt"
	"os"
	"vpn-dns/internal/vpndns"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/process"

	"github.com/spf13/cobra"
)

// AppName represents app name.
const AppName = "vpn-dns"

// PackageName represents app package name.
const PackageName = "co.myrt.vpndns"

// Version represents current app version.
var Version = "development"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     AppName,
	Version: Version,
	Short:   "An app that fixes macOS DNS behavior when using a VPN",
}

// Execute is the main CLI entrypoint.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var configPath string

func init() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	defaultPath := userDir + "/.config/" + AppName + "/config.yaml"
	rootCmd.PersistentFlags().StringVarP(
		&configPath,
		"config", "c",
		defaultPath,
		"Configuration file path")
}

func createApp() (vpndns.Changer, process.Daemon) {
	app, err := vpndns.NewChanger(configPath, exec.Run)
	if err != nil {
		fmt.Println("Can't initialize app:", err.Error())
		os.Exit(1)
	}
	daemon := process.NewDaemon(AppName)
	return app, daemon
}
