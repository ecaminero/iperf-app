package commands

import (
	"fmt"
	"runtime"
	"time"
)

type IperfCommands struct{}

func NewIperfCommands() *IperfCommands {
	return &IperfCommands{}
}

func (ic *IperfCommands) GetIperfInternetNationalCommand(title, duration, parallel string) []string {
	timestamp := time.Now().Format("2006-01-02")
	fmt.Printf("Prueba iniciada a las %s", timestamp)

	if runtime.GOOS == "windows" {
		return []string{"cd C:\\iperf-3.1.3-win64",
			fmt.Sprintf("iperf3 -c 181.212.30.100 -t %s -P %s --get-server-output --logfile INET_NAC_%s_%s.txt -T %s_%s",
				duration, parallel, timestamp, title, timestamp, title),
			fmt.Sprintf("echo Prueba terminada a las %%time%% del %%date%%"),
			"start %windir%\\explorer.exe \"C:\\iperf-3.1.3-win64\"",
		}
	}

	// Linux/Unix
	return []string{
		fmt.Sprintf("iperf3 -c 181.212.30.100 -t %s -P %s --get-server-output --logfile INET_NAC_%s_%s.txt -T %s_%s",
			duration, parallel, timestamp, title, timestamp, title),
		fmt.Sprintf("echo 'Prueba terminada a las' $(date)"),
	}
}

// GetIperfInternetInternationalCommand retorna el comando para test Internet Internacional
func (ic *IperfCommands) GetIperfInternetInternationalCommand(title, duration, parallel string) []string {
	timestamp := time.Now().Format("2006-01-02")

	if runtime.GOOS == "windows" {
		return []string{
			"cd C:\\iperf-3.1.3-win64",
			fmt.Sprintf("echo Prueba iniciada a las %%time%% del %%date%%"),
			fmt.Sprintf("iperf3 -c iperf.volia.net -t %s -P %s --get-server-output --logfile INET_INTERNAC_%s_%s.txt -T %s_%s",
				duration, parallel, timestamp, title, timestamp, title),
			fmt.Sprintf("echo Prueba terminada a las %%time%% del %%date%%"),
			"start %windir%\\explorer.exe \"C:\\iperf-3.1.3-win64\"",
		}
	}

	// Linux/Unix
	return []string{
		fmt.Sprintf("echo 'Prueba iniciada a las' $(date)"),
		fmt.Sprintf("iperf3 -c iperf.volia.net -t %s -P %s --get-server-output --logfile INET_INTERNAC_%s_%s.txt -T %s_%s",
			duration, parallel, timestamp, title, timestamp, title),
		fmt.Sprintf("echo 'Prueba terminada a las' $(date)"),
	}
}

// GetIperfMPLSCommand retorna el comando para test MPLS
func (ic *IperfCommands) GetIperfMPLSCommand(title, ipserver, duration, parallel string) []string {
	timestamp := time.Now().Format("2006-01-02")

	if runtime.GOOS == "windows" {
		return []string{
			"cd C:\\iperf-3.1.3-win64",
			fmt.Sprintf("echo Prueba iniciada a las %%time%% del %%date%%"),
			fmt.Sprintf("iperf3 -c %s -t %s -P %s --get-server-output --logfile MPLS_%s_%s.txt -T %s_%s",
				ipserver, duration, parallel, timestamp, title, timestamp, title),
			fmt.Sprintf("echo Prueba terminada a las %%time%% del %%date%%"),
			"start %windir%\\explorer.exe \"C:\\iperf-3.1.3-win64\"",
		}
	}

	// Linux/Unix
	return []string{
		fmt.Sprintf("echo 'Prueba iniciada a las' $(date)"),
		fmt.Sprintf("iperf3 -c %s -t %s -P %s --get-server-output --logfile MPLS_%s_%s.txt -T %s_%s",
			ipserver, duration, parallel, timestamp, title, timestamp, title),
		fmt.Sprintf("echo 'Prueba terminada a las' $(date)"),
	}
}

func (ic *IperfCommands) GetIperfP2PClientCommand(title, ip, duration, parallel string) []string {
	timestamp := time.Now().Format("2006-01-02")

	if runtime.GOOS == "windows" {
		return []string{
			"cd C:\\iperf-3.1.3-win64",
			fmt.Sprintf("echo Prueba iniciada a las %%time%% del %%date%%"),
			fmt.Sprintf("iperf3 -c %s -t %s -P %s --get-server-output --logfile P2PL2_%s_%s.txt -T %s_%s",
				ip, duration, parallel, timestamp, title, timestamp, title),
			fmt.Sprintf("echo Prueba terminada a las %%time%% del %%date%%"),
			"start %windir%\\explorer.exe \"C:\\iperf-3.1.3-win64\"",
		}
	}

	// Linux/Unix
	return []string{
		fmt.Sprintf("echo 'Prueba iniciada a las' $(date)"),
		fmt.Sprintf("iperf3 -c %s -t %s -P %s --get-server-output --logfile P2PL2_%s_%s.txt -T %s_%s",
			ip, duration, parallel, timestamp, title, timestamp, title),
		fmt.Sprintf("echo 'Prueba terminada a las' $(date)"),
	}
}

// GetIperfP2PServerCommand retorna el comando para test P2P Servidor
func (ic *IperfCommands) GetIperfP2PServerCommand() []string {
	if runtime.GOOS == "windows" {
		return []string{
			"cd C:\\iperf-3.1.3-win64",
			"iperf3 -s",
		}
	}

	// Linux/Unix
	return []string{
		"iperf3 -s",
	}
}
