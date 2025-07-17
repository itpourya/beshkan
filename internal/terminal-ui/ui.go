package terminalui

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/itpourya/beshkan/internal/configs"
	dnsreader "github.com/itpourya/beshkan/internal/dnsReader"
)

func ShowTerminalUserInterface() ([]string, error) {
	printHeader("Current DNS:")
	cdns, err := dnsreader.GetCurrentDnsServerAddr()
	if err != nil {
		return nil, err
	}
	currentDNS(cdns)

	printHeader("Available DNS Providers:")
	for i, provider := range configs.Providers {
		color.Cyan("%d) %s\t%v", i+1, provider.Name, provider.IPs)
	}
	color.Cyan("0) Exit")

	var choice int
	color.Magenta("\nSelect (0-%d): ", len(configs.Providers))
	fmt.Scan(&choice)

	if choice == 0 {
		color.Green("Bye!")
		return []string{}, nil
	}
	if choice < 1 || choice > len(configs.Providers) {
		color.Red("Invalid choice!")
		return []string{}, nil
	}

	selected := configs.Providers[choice-1]

	return selected.IPs, nil
}

func currentDNS(dnsAddr [][]string) {
	if dnsAddr == nil || len(dnsAddr) == 0 {
		color.Blue("NOTHING")
		return
	}
	color.Blue("• %s %s", dnsAddr[0][1], dnsAddr[1][1])
}

func printHeader(title string) {
	sep := "─"
	color.New(color.FgHiWhite, color.Bold).Println("\n" + title)
	color.New(color.FgHiBlack).Println(sep + sep + sep + sep + sep + sep + sep + sep + sep + sep + sep + sep + sep + sep + sep)
}
