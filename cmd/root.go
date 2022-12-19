package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Version represents current app version.
var Version = "development"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "vpn-dns",
	Version: Version,
	Short:   "An app that fixes macOS DNS behavior when using a VPN",
}

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
	defaultPath := userDir + "/.config/vpn-dns/config.yaml"
	rootCmd.PersistentFlags().StringVarP(
		&configPath,
		"config", "c",
		defaultPath,
		"Configuration file path")
}
