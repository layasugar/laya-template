package gcnf

import (
	"log"
	"net"

	"github.com/layasugar/laya/core/constants"
)

var ip = ""

func LocalIP() string {
	if ip == "" {
		ips := getLocalIPs()
		if len(ips) > 0 {
			ip = ips[0] + ":80"
		} else {
			ip = constants.DEFAULT_LISTEN
		}
	}

	return ip
}

func getLocalIPs() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
