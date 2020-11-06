package util

import (
	"net"
	"os"
)

func GetHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "Unknown-host"
	}

	return hostName
}

func GetHostIP() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "0.0.0.0"
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return "0.0.0.0"
}
