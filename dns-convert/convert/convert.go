package convert

import (
	"context"
	"log"
	"net"
	"sort"
)

func ConvertNames(entries []string) []string {
	var ipList []string

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", net.JoinHostPort("1.1.1.1", "53"))
		},
	}

	for _, name := range entries {
		byteAddress, err := r.LookupIPAddr(context.Background(), name)

		if err != nil {
			log.Fatal(err)
		}

		for _, ip := range byteAddress {
			ipList = append(ipList, name+","+ip.String()+"\n")
		}
	}
	sort.Strings(ipList)
	return ipList
}
