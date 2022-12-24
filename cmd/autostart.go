package cmd

import (
	"fmt"
	"os"
	"vpn-dns/pkg/login"

	"github.com/spf13/cobra"
)

func createLoginItem() login.Item {
	const binPath = "/usr/local/bin/" + AppName
	itemPath, err := login.LaunchAgentPath(PackageName)
	if err != nil {
		fmt.Println("Can't initialize app:", err.Error())
		os.Exit(1)
	}
	item := login.NewItem(
		PackageName,
		binPath+" -c "+configPath+" start",
		itemPath,
	)
	return item
}

// autostartEnableCmd represents the `autostart enable` command.
var autostartEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables automatic startup",
	Run: func(cmd *cobra.Command, args []string) {
		item := createLoginItem()
		err := item.Write()
		if item.IsSet() {
			fmt.Println("Autostart is already enabled.")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("Can't write login item:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Autostart is enabled.")
	},
}

// autostartDisableCmd represents the `autostart disable` command.
var autostartDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables automatic startup",
	Run: func(cmd *cobra.Command, args []string) {
		item := createLoginItem()
		if !item.IsSet() {
			fmt.Println("Autostart is not enabled.")
			os.Exit(1)
		}
		err := item.Remove()
		if err != nil {
			fmt.Println("Can't remove login item:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Autostart is disabled.")
	},
}

// autostartCmd represents the autostart command.
var autostartCmd = &cobra.Command{
	Use:   "autostart",
	Short: "Enables or disables automatic startup",
	Run: func(cmd *cobra.Command, args []string) {
		item := createLoginItem()
		if item.IsSet() {
			fmt.Println("Autostart is enabled.")
			fmt.Println("To disable, run: vpn-dns autostart disable")
		} else {
			fmt.Println("Autostart is not enabled.")
			fmt.Println("To enable, run: vpn-dns autostart enable")
		}
	},
}

func init() {
	autostartCmd.AddCommand(autostartEnableCmd)
	autostartCmd.AddCommand(autostartDisableCmd)
	rootCmd.AddCommand(autostartCmd)
}
