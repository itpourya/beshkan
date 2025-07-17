package dnschanger

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/itpourya/beshkan/internal/notification"
)

func ChangeCurrentDnsServer(dnsAddrs string) error {
	  conName, err := getActiveConnection()
    if err != nil {
				return err
    }

    fmt.Println("🔌 Active Connection", conName)

    cmds := [][]string{
        {"nmcli", "con", "mod", conName, "ipv4.dns", dnsAddrs},
        {"nmcli", "con", "mod", conName, "ipv4.ignore-auto-dns", "yes"},
        {"nmcli", "con", "up", conName},
    }

    for _, cmdArgs := range cmds {
        if err := runCommand(cmdArgs); err != nil {
            return err
        }
    }

    fmt.Println("✅ DNS CHANGED TO", dnsAddrs)
		notification.SendNotification("Beshkan", "Dns changed to "+dnsAddrs)
		return nil
}

func getActiveConnection() (string, error) {
    cmd := exec.Command("nmcli", "-t", "-f", "NAME,DEVICE", "con", "show", "--active")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return "", err
    }

    lines := strings.Split(out.String(), "\n")
    for _, line := range lines {
        if strings.TrimSpace(line) == "" {
            continue
        }
        parts := strings.Split(line, ":")
        if len(parts) >= 2 {
            return parts[0], nil
        }
    }

    return "", fmt.Errorf("can't find active connection")
}

func runCommand(args []string) error {
    cmd := exec.Command(args[0], args[1:]...)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("%s: %s", err, stderr.String())
    }
    return nil
}
