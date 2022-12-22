package cmd

import (
	"fmt"
	"vpn-dns/internal/login"

	"github.com/spf13/cobra"
)

// autostartCmd represents the autostart command.
var autostartCmd = &cobra.Command{
	Use:   "autostart",
	Short: "Enables or disables automatic startup",
	Run: func(cmd *cobra.Command, args []string) {
		packageName := "co.myrt.vpndns"
		binPath := "/usr/local/bin/vpn-dns"
		itemPath, err := login.LaunchAgentPath(packageName)
		if err != nil {
			panic(err)
		}
		item := login.NewItem(
			packageName,
			binPath+" -c "+configPath+" start",
			itemPath,
		)
		if err != nil {
			return
		}
		isSet := item.IsSet()
		switch len(args) {
		case 0:
			if isSet {
				fmt.Println("Autostart is enabled.")
				fmt.Println("To disable, run: vpn-dns autostart disable")
			} else {
				fmt.Println("Autostart is not enabled.")
				fmt.Println("To enable, run: vpn-dns autostart enable")
			}
		case 1:
			switch args[0] {
			case "enable":
				err = item.Write()
				fmt.Println("Enabled.")
			case "disable":
				err = item.Remove()
				fmt.Println("Disabled.")
			default:
				fmt.Println("Unknown command.")
			}
		default:
			fmt.Println("Unknown command.")
		}
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(autostartCmd)
}
