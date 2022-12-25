// VPN DNS Changer.
package main

import (
	"net/http"
	_ "net/http/pprof"
	"vpn-dns/cmd"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	cmd.Execute()
}
