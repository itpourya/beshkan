package main

import (
	"log"
	"strings"

	dnschanger "github.com/itpourya/beshkan/internal/dnsChanger"
	terminalui "github.com/itpourya/beshkan/internal/terminal-ui"
)

func main() {
	ipv4, err := terminalui.ShowTerminalUserInterface()
	if err != nil {
		log.Println("Show terminal interface error :", err)
	}
	if ipv4 == nil {
		return
	}

	dnsaddr := strings.Join(ipv4, " ")
	err = dnschanger.ChangeCurrentDnsServer(dnsaddr)
	if err != nil {
		log.Println("Change DNS error :", err)
	}
}
