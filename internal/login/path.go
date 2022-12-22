package login

import "os"

func LaunchAgentPath(name string) (string, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return userDir + "/Library/LaunchAgents/" + name + ".plist", nil
}
