package main

import (
	"net"
	"os"
)

func GetIP() (ip []string) {
	host, _ := os.Hostname()
	addrs, err := net.LookupIP(host)
	if err != nil {
		return []string{}
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			if !contains(ip, ipv4.String()) {
				ip = append(ip, ipv4.String())
			}
		}
	}
	return ip
}

func contains(s []string, v string) bool {
	for _, value := range s {
		if value == v {
			return true
		}
	}
	return false
}
