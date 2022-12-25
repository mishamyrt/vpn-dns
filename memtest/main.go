// A special version of the application which includes additional functions for memory debugging.
// For details, see the memtest section of the Makefile.
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"vpn-dns/cmd"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	}()
	cmd.Execute()
}
