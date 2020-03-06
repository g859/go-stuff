package unique

import (
	"strings"
)

func UniqueIPs(ipList []string) []string {
	if len(ipList) == 0 {
		return ipList
	}
	seen := make([]string, 0, len(ipList))
slice:
	for i, n := range ipList {
		if i == 0 {
			ipList = ipList[:0]
		}
		for _, t := range seen {
			if n == t {
				continue slice
			}
		}
		seen = append(seen, n)
		ipList = append(ipList, strings.TrimSpace(n)+"\n")
	}
	return ipList
}
