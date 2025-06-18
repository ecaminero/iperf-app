package entities

// NetworkConfig contiene la configuración de red para ZTP
type NetworkConfig struct {
	InterfaceName string // "Ethernet" en Windows, "eth0" en Linux
	IPAddress     string // Ej: "10.255.30.38"
	SubnetMask    string // Ej: "255.255.255.252"
	Gateway       string // Ej: "10.255.30.37"
}
