package notification

import "os/exec"

func SendNotification(title, message string) {
	exec.Command("notify-send", title, message).Run()
}
