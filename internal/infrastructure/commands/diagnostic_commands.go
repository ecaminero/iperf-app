package commands

import (
	"fmt"
	"runtime"
)

type DiagnosticCommands struct{}

func NewDiagnosticCommands() *DiagnosticCommands {
	return &DiagnosticCommands{}
}

// GetPingCommand retorna el comando de ping
func (dc *DiagnosticCommands) GetPingCommand(target string, count int) []string {
	if count <= 0 {
		count = 4 // default Windows
	}

	if runtime.GOOS == "windows" {
		return []string{
			fmt.Sprintf("ping -n %d %s", count, target),
		}
	}

	return []string{
		fmt.Sprintf("ping -c %d %s", count, target),
	}
}

// GetTracerouteCommand retorna el comando de traceroute
func (dc *DiagnosticCommands) GetTracerouteCommand(target string) []string {
	if runtime.GOOS == "windows" {
		return []string{
			fmt.Sprintf("tracert -d %s", target),
		}
	}

	return []string{
		fmt.Sprintf("traceroute -n %s || tracepath %s", target, target),
	}
}

// GetRouteTableCommand retorna el comando para mostrar tabla de rutas
func (dc *DiagnosticCommands) GetRouteTableCommand() []string {
	if runtime.GOOS == "windows" {
		return []string{
			"route print",
		}
	}

	return []string{
		"ip route show || route -n",
	}
}

// GetAddRouteCommand retorna el comando para agregar ruta estática
func (dc *DiagnosticCommands) GetAddRouteCommand(destination, mask, gateway string) []string {
	if runtime.GOOS == "windows" {
		return []string{
			fmt.Sprintf("route add %s mask %s %s", destination, mask, gateway),
		}
	}

	// Linux/Unix - usando comando moderno (ip)
	nc := &NetworkCommands{}
	cidr := nc.getCIDRFromMask(mask)
	return []string{
		fmt.Sprintf("ip route add %s/%s via %s", destination, cidr, gateway),
	}
}

// GetAddRouteCommandLegacy retorna el comando legacy para agregar ruta
func (dc *DiagnosticCommands) GetAddRouteCommandLegacy(destination, mask, gateway string) []string {
	if runtime.GOOS == "windows" {
		return dc.GetAddRouteCommand(destination, mask, gateway)
	}

	// Linux/Unix - usando comando legacy (route)
	return []string{
		fmt.Sprintf("route add -net %s netmask %s gw %s", destination, mask, gateway),
	}
}

// GetOpenFolderCommand retorna el comando para abrir carpeta
func (dc *DiagnosticCommands) GetOpenFolderCommand(path string) []string {
	if runtime.GOOS == "windows" {
		if path == "" {
			path = "C:\\iperf-3.1.3-win64"
		}
		return []string{
			fmt.Sprintf("start %%windir%%\\explorer.exe \"%s\"", path),
		}
	}

	// Linux/Unix
	if path == "" {
		path = "/usr/local/bin" // or current directory
	}

	// Try different file managers
	return []string{
		fmt.Sprintf("xdg-open %s || nautilus %s || dolphin %s || thunar %s || pcmanfm %s",
			path, path, path, path, path),
	}
}
