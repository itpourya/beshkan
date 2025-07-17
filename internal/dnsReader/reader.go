package dnsreader

import (
	"os"
	"regexp"
)

func GetCurrentDnsServerAddr() ([][]string, error) {
	data, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}

	dnsAddr := validate(string(data))

	return dnsAddr, nil
}

func validate(dnsAddr string) [][]string {
	pattern := "nameserver\\s+(\\d{1,3}(?:\\.\\d{1,3}){3})"
	re := regexp.MustCompile(pattern)
	maches := re.FindAllStringSubmatch(dnsAddr, -1)

	return maches
}
