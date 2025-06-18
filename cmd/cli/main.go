package main

import (
	"bufio"
	"fmt"
	"iperf-app/internal/infrastructure/cli/ui"
	"iperf-app/internal/infrastructure/executors"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandExecutor := executors.NewCommandExecutor()

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
			showDiagnosticMenu(commandExecutor)
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

func showDiagnosticMenu(commandExecutor *executors.CommandExecutor) {
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
			handleDiagPing(commandExecutor)
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

func handleDiagPing(commandExecutor *executors.CommandExecutor) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Indicar la direccion IP del servidor (predeterminado 10.255.30.3): ")
	scanner.Scan()
	ipserver := strings.TrimSpace(scanner.Text())
	if ipserver == "" {
		ipserver = "10.255.30.3"
	}

	count := 4 // valor por defecto
	result, err := commandExecutor.ExecutePing(ipserver, count)
	if err != nil {
		fmt.Printf("Error ejecutando ping: %v\n", err)
		if result != nil && result.ErrorOutput != "" {
			fmt.Printf("Detalles del error:\n%s\n", result.ErrorOutput)
		}
	} else {
		fmt.Println("Ping completado exitosamente")
		if result.Output != "" {
			fmt.Printf("Resultado:\n%s\n", result.Output)
		}
	}

	fmt.Println()
	fmt.Print("Presione Enter para continuar...")
	scanner.Scan()
}
