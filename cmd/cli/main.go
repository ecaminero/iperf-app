package main

import (
	"bufio"
	"fmt"
	"iperf-app/internal/infrastructure/cli/ui"
	"os"
	"strings"
)

func main() {
	fmt.Print(" _                  __\n(_)_ __   ___ _ __ / _|   __ _ _ __  _ __\n| | '_ \\ / _ \\ '__| |_   / _` | '_ \\| '_ \\\n| | |_) |  __/ |  |  _| | (_| | |_) | |_) |\n|_| .__/ \\___|_|  |_|    \\__,_| .__/| .__/\n  |_|                         |_|   |_|")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		ui.MainMenu()
		fmt.Print("Su opcion es: ")
		if !scanner.Scan() {
			break
		}
		option := strings.TrimSpace(scanner.Text())
		switch option {
		case "1":
			fmt.Println("Hello ZTP")
		case "2":
			showIperfInternetMenu()
		case "3":
			fmt.Println("handleIperfMPLS()")
		case "4":
			showIperfP2PMenu()
		case "5":
			showDiagnosticMenu()
		case "6":
			fmt.Println("handleOpenFolder()")
		case "7":
			fmt.Println("Saliendo...")
			return
		case "":
			// Si no ingresa nada, volver a mostrar el menú
			continue
		default:
			fmt.Printf("Opción inválida: %s\n", option)
			fmt.Println("Presione Enter para continuar...")
			scanner.Scan()
		}
	}
}

func showIperfInternetMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[2J\033[H")
		ui.IperfInternetMenu()

		fmt.Print("Su opcion es: ")
		if !scanner.Scan() {
			return
		}

		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			fmt.Println("handleIperfInternetNational")
		case "2":
			fmt.Println("handleIperfInternetInternational()")
		case "3":
			return // Volver al menú principal
		case "":
			continue
		default:
			fmt.Printf("Opción inválida: %s\n", option)
			fmt.Println("Presione Enter para continuar...")
			scanner.Scan()
		}
	}
}

func showIperfP2PMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[2J\033[H")
		ui.IperfP2PMenu()

		fmt.Print("Su opcion es: ")
		if !scanner.Scan() {
			return
		}

		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			fmt.Println("handleIperfInternetNational")
		case "2":
			fmt.Println("handleIperfInternetInternational()")
		case "3":
			return // Volver al menú principal
		case "":
			continue
		default:
			fmt.Printf("Opción inválida: %s\n", option)
			fmt.Println("Presione Enter para continuar...")
			scanner.Scan()
		}
	}
}

func showDiagnosticMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[2J\033[H")
		ui.DiagnosticMenu()

		fmt.Print("Su opcion es: ")
		if !scanner.Scan() {
			return
		}

		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			fmt.Println("handleDiagPing()")

		case "2":
			fmt.Println("handleDiagTraceroute()")
		case "3":
			fmt.Println("handleDiagRouteTable()")
		case "4":
			fmt.Println("handleDiagAddRouteMPLS()")
		case "5":
			fmt.Println("handleDiagAddRouteP2P()")
		case "6":
			return // Volver al menú principal
		case "":
			continue
		default:
			fmt.Printf("Opción inválida: %s\n", option)
			fmt.Println("Presione Enter para continuar...")
			scanner.Scan()
		}
	}
}
