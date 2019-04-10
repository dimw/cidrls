package cidr

import (
	"net"
	"strings"
)

// Based on https://gist.github.com/kotakanbe/d3059af990252ba89a82
func Hosts(cidr string) ([]string, error) {
	if !strings.Contains(cidr, "/") {
		cidr = cidr + "/32"
	}

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func removeNetworkAndBroadcast(ips []string) []string {
	return ips[1 : len(ips)-1]
}

func CalculateIps(cidr string, removeNetworkAndBroadcastFlag bool) []string {
	ips, _ := Hosts(cidr)

	if removeNetworkAndBroadcastFlag && len(ips) > 2 {
		ips = removeNetworkAndBroadcast(ips)
	}

	return ips
}
