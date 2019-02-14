package main

import (
	"fmt"
	"net"
	"strings"
)

func isGoogle(name string) bool {
	if strings.HasSuffix(name, "googlebot.com.") || strings.HasSuffix(name, "google.com.") {
		return true
	}
	return false
}

func googleBotCheck(addr string) bool {
	names, err := net.LookupAddr(addr)
	if err != nil {
		return false
	}

	for _, name := range names {
		if !isGoogle(name) {
			continue
		}
		ips, err := net.LookupIP(name)
		if err != nil {
			continue
		}

		for _, ip := range ips {
			if ip.String() == addr {
				return true
			}
		}
	}
	return false
}

const (
	googleBotA = "66.249.66.1"
	googleBotB = "66.249.90.77"
	gcp        = "74.125.0.1"
)

func main() {
	fmt.Println(googleBotCheck(googleBotA))
	fmt.Println(googleBotCheck(googleBotB))
	fmt.Println(googleBotCheck(gcp))
}
