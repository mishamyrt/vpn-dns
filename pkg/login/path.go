// Package login provides tools for working with autorun on macOS.
package login

import "os"

// LaunchAgentPath generates path for current user's package login item.
func LaunchAgentPath(name string) (string, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return userDir + "/Library/LaunchAgents/" + name + ".plist", nil
}
