package commands

import (
	"fmt"
	"runtime"
)

// NetworkCommands contiene todos los comandos de red específicos por OS
type NetworkCommands struct{}

func NewNetworkCommands() *NetworkCommands {
	return &NetworkCommands{}
}

// GetZTPCommand retorna el comando de autoaprovisionamiento (ZTP)
func (nc *NetworkCommands) GetZTPCommand(interfaceName, ipAddress, subnetMask, gateway string) []string {
	if runtime.GOOS == "windows" {
		return []string{
			fmt.Sprintf("netsh interface ipv4 set address name=\"%s\" static %s %s %s",
				interfaceName, ipAddress, subnetMask, gateway),
		}
	}

	// Linux/Unix - usando comandos modernos (ip)
	cidr := nc.getCIDRFromMask(subnetMask)
	return []string{
		fmt.Sprintf("ip addr flush dev %s", interfaceName),
		fmt.Sprintf("ip addr add %s/%s dev %s", ipAddress, cidr, interfaceName),
		fmt.Sprintf("ip link set %s up", interfaceName),
		fmt.Sprintf("ip route add default via %s", gateway),
	}
}

// GetZTPCommandLegacy retorna el comando ZTP usando herramientas legacy en Linux
func (nc *NetworkCommands) GetZTPCommandLegacy(interfaceName, ipAddress, subnetMask, gateway string) []string {
	if runtime.GOOS == "windows" {
		return nc.GetZTPCommand(interfaceName, ipAddress, subnetMask, gateway)
	}

	// Linux/Unix - usando comandos legacy (ifconfig/route)
	return []string{
		fmt.Sprintf("ifconfig %s %s netmask %s", interfaceName, ipAddress, subnetMask),
		fmt.Sprintf("route add default gw %s %s", gateway, interfaceName),
	}
}

// getCIDRFromMask convierte máscara de subred a notación CIDR
func (nc *NetworkCommands) getCIDRFromMask(mask string) string {
	maskMap := map[string]string{
		"255.255.255.255": "32",
		"255.255.255.252": "30",
		"255.255.255.248": "29",
		"255.255.255.240": "28",
		"255.255.255.224": "27",
		"255.255.255.192": "26",
		"255.255.255.128": "25",
		"255.255.255.0":   "24",
		"255.255.254.0":   "23",
		"255.255.252.0":   "22",
		"255.255.248.0":   "21",
		"255.255.240.0":   "20",
		"255.255.224.0":   "19",
		"255.255.192.0":   "18",
		"255.255.128.0":   "17",
		"255.255.0.0":     "16",
		"255.254.0.0":     "15",
		"255.252.0.0":     "14",
		"255.248.0.0":     "13",
		"255.240.0.0":     "12",
		"255.224.0.0":     "11",
		"255.192.0.0":     "10",
		"255.128.0.0":     "9",
		"255.0.0.0":       "8",
	}

	if cidr, exists := maskMap[mask]; exists {
		return cidr
	}

	return "24" // default
}
